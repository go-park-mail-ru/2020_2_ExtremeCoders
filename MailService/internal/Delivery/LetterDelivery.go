package Delivery

import (
	"MailService/internal/UseCase"
	"MailService/pkg/convert"
	pb "MailService/proto"
)

type Interface interface {
	GetLettersByDir(dir *pb.DirName) *pb.LetterListResponse
	SaveLetter(letter *pb.Letter) *pb.Response
	WatchedLetter(Lid *pb.Lid)  *pb.Response
}

type Delivery struct{
	uc UseCase.UseCase
}

func (ld Delivery)GetLettersByDir(dir *pb.DirName) *pb.LetterListResponse{
	err, Letters:=ld.uc.GetLettersByDir(dir.DirName)
	LetterResp:=pb.LetterListResponse{}
	resp:=pb.Response{Ok: false, Description: err.Error()}
	if err==nil{
		resp.Ok=true
	}
	for _, letter:=range Letters{
		pbLetter:=convert.ModelToProto(letter)
		LetterResp.Letter=append(LetterResp.Letter, &pbLetter)
	}
	LetterResp.Result=&resp
	return &LetterResp
}

func (ld Delivery)SaveLetter(letter *pb.Letter) *pb.Response{
	letter.IsWatched=false;
	err:=ld.uc.SaveLetter(convert.ProtoToModel(letter))

	resp:=pb.Response{Ok: false, Description: err.Error()}
	if err==nil{
		resp.Ok=true
	}
	return &resp
}

func (ld Delivery)WatchedLetter(Lid *pb.Lid)  *pb.LetterResponse{
	err, letter:=ld.uc.WatchLetter(Lid.Lid)
	resp:=pb.Response{Ok: false, Description: err.Error()}
	if err==nil{
		resp.Ok=true
	}
	letterPB:=convert.ModelToProto(letter)
	return &pb.LetterResponse{Letter: &letterPB, Result: &resp}
}
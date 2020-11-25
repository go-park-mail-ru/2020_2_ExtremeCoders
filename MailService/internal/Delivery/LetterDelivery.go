package Delivery

import (
	"MailService/internal/UseCase"
	"MailService/pkg/convert"
	pb "MailService/proto"
)

type Interface interface {
	GetLettersByDirRecv(dir *pb.DirName) pb.LetterListResponse
	GetLettersByDirSend(dir *pb.DirName) pb.LetterListResponse
	SaveLetter(letter *pb.Letter) *pb.Response
	WatchedLetter(Lid *pb.Lid)  *pb.Response
}

type Delivery struct{
	uc UseCase.UseCase
}

func (ld Delivery)GetLettersByDirRecv(dir *pb.DirName) *pb.LetterListResponse{
	err, letters:=ld.uc.GetLettersRecv(dir.DirName)
	resp:=pb.Response{Ok: false, Description: err.Error()}
	if err==nil{
		resp.Ok=true
	}
	lettersListPb:=convert.ModelToProtoList(&letters)
	letterPb:=pb.LetterListResponse{Result: &resp, Letter: lettersListPb}
	return &letterPb
}

func (ld Delivery)GetLettersByDirSend(dir *pb.DirName) *pb.LetterListResponse{
	err, letters:=ld.uc.GetLettersSend(dir.DirName)
	resp:=pb.Response{Ok: false, Description: err.Error()}
	if err==nil{
		resp.Ok=true
	}
	lettersListPb:=convert.ModelToProtoList(&letters)
	letterPb:=pb.LetterListResponse{Result: &resp, Letter: lettersListPb}
	return &letterPb
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
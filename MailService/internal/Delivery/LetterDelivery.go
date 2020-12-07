package Delivery

import (
	"MailService/internal/UseCase"
	"MailService/pkg/convert"
	pb "MailService/proto"
	"context"
	"fmt"
)

type Delivery struct {
	uc UseCase.Interface
}

func New(usecase UseCase.Interface) pb.LetterServiceServer {
	return Delivery{uc: usecase}
}

func (ld Delivery) GetLettersByDirRecv(ctx context.Context, dir *pb.DirName) (*pb.LetterListResponse, error) {
	fmt.Println("recv dir")
	err, letters := ld.uc.GetLettersRecvDir(dir.DirName, dir.Limit, dir.Offset)
	resp := pb.Response{Ok: true}
	if err != nil || letters == nil {
		return &pb.LetterListResponse{}, err
	}
	lettersListPb := convert.ModelToProtoList(&letters)
	letterPb := pb.LetterListResponse{Result: &resp, Letter: lettersListPb}
	return &letterPb, nil
}

func (ld Delivery) GetLettersByDirSend(ctx context.Context, dir *pb.DirName) (*pb.LetterListResponse, error) {
	err, letters := ld.uc.GetLettersSendDir(dir.DirName)
	resp := pb.Response{Ok: true}
	if err != nil || letters == nil {
		return &pb.LetterListResponse{}, err
	}
	lettersListPb := convert.ModelToProtoList(&letters)
	letterPb := pb.LetterListResponse{Result: &resp, Letter: lettersListPb}
	return &letterPb, nil
}

func (ld Delivery) SaveLetter(ctx context.Context, letter *pb.Letter) (*pb.Response, error) {
	letter.IsWatched = false
	err := ld.uc.SaveLetter(convert.ProtoToModel(letter))
	resp := pb.Response{Ok: true, Description: ""}
	if err != nil {
		resp.Ok = false
		resp.Description = err.Error()
	}
	return &resp, nil
}

func (ld Delivery) WatchedLetter(ctx context.Context, Lid *pb.Lid) (*pb.LetterResponse, error) {
	err, letter := ld.uc.WatchLetter(Lid.Lid)
	resp := pb.Response{Ok: true, Description: ""}
	if err != nil {
		resp.Ok = false
		resp.Description = err.Error()
	}
	letterPB := convert.ModelToProto(letter)
	return &pb.LetterResponse{Letter: &letterPB, Result: &resp}, nil
}

func (ld Delivery) GetLettersRecv(ctx context.Context, email *pb.Email) (*pb.LetterListResponse, error) {
	err, letters := ld.uc.GetLettersRecv(email.Email, email.Limit, email.Offset)
	resp := pb.Response{Ok: true}
	if err != nil || letters == nil {
		return &pb.LetterListResponse{}, err
	}
	lettersListPb := convert.ModelToProtoList(&letters)
	letterPb := pb.LetterListResponse{Result: &resp, Letter: lettersListPb}
	return &letterPb, nil
}

func (ld Delivery) GetLettersSend(ctx context.Context, email *pb.Email) (*pb.LetterListResponse, error) {
	err, letters := ld.uc.GetLettersSend(email.Email)
	resp := pb.Response{Ok: true}
	if err != nil || letters == nil {
		return &pb.LetterListResponse{}, err
	}
	lettersListPb := convert.ModelToProtoList(&letters)
	letterPb := pb.LetterListResponse{Result: &resp, Letter: lettersListPb}
	return &letterPb, nil
}

func (ld Delivery) AddLetterToDir(ctx context.Context, dirlid *pb.DirLid) (*pb.Response, error) {
	err := ld.uc.AddLetterToDir(dirlid.Lid, dirlid.Did, dirlid.Type)

	resp := pb.Response{Ok: true, Description: ""}
	if err != nil {
		resp.Ok = false
		resp.Description = err.Error()
	}
	return &resp, nil
}

func (ld Delivery) RemoveLetterFromDir(ctx context.Context, dirlid *pb.DirLid) (*pb.Response, error) {
	err := ld.uc.RemoveLetterFromDir(dirlid.Lid, dirlid.Did, dirlid.Type)

	resp := pb.Response{Ok: true, Description: ""}
	if err != nil {
		resp.Ok = false
		resp.Description = err.Error()
	}
	return &resp, nil
}

func (ld Delivery) RemoveDir(ctx context.Context, dirlid *pb.DirLid) (*pb.Response, error) {
	err := ld.uc.RemoveDir(dirlid.Did, dirlid.Type)
	resp := pb.Response{Ok: true, Description: ""}
	if err != nil {
		resp.Ok = false
		resp.Description = err.Error()
	}
	return &resp, nil
}

func (ld Delivery) RemoveLetter(ctx context.Context, Lid *pb.Lid)(*pb.Response, error){
	err:=ld.uc.RemoveLetter(Lid.Lid)
	resp := pb.Response{Ok: true, Description: ""}
	if err!=nil{
		resp.Ok=false
		resp.Description=err.Error()
	}else{
		resp.Description="ok"
	}
	return &resp, nil
}

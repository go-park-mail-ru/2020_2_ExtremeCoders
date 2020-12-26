package FolderDelivery

import (
	Model "Mailer/MainApplication/internal/Letter/LetterModel"
	mailProto "Mailer/MailService/proto"
	userProto "Mailer/UserService/proto"
	"encoding/json"
)

func ProtoFolderListResponse(folders []*userProto.FolderNameType) []byte {
	ans := FolderList{
		Code:    200,
		Folders: ProtoToModelList(folders),
	}
	res, err := json.Marshal(ans)
	if err!=nil{
		return nil
	}
	return res
}

func SuccessRespAns() []byte {
	ans := SuccessAns{Code: 200}
	res, _ := json.Marshal(ans)
	return res
}

func GetFoldersError(err error) []byte {
	ans := &LetterErr{
		Code:        400,
		Description: err.Error(),
	}
	res, _ := json.Marshal(ans)
	return res
}

func ProtoResponseAnswer(pbLetter *mailProto.Response) []byte {
	code := 200
	if pbLetter==nil{
		ans := LetterList{
			Code:        500,
		}
		res, _ := json.Marshal(ans)
		return res
	}
	if !pbLetter.Ok {
		code = 409
	}
	ans := LetterErr{
		Code:        code,
		Description: pbLetter.Description,
	}
	res, _ := json.Marshal(ans)
	return res
}

func ProtoLetterListAnswer(pbLetter *mailProto.LetterListResponse) []byte {
	code := 200
	if pbLetter==nil || pbLetter.Letter==nil || pbLetter.Result==nil{
		ans := LetterList{
			Code:        500,
		}
		res, _ := json.Marshal(ans)
		return res
	}
	if !pbLetter.Result.Ok {
		code = 409
	}
	ans := LetterList{
		Code:        code,
		Description: pbLetter.Result.Description,
		Letter:      ProtoToModelMail(pbLetter),
	}
	res, _ := json.Marshal(ans)
	return res
}

func ProtoToModelList(pbLetter []*userProto.FolderNameType) []Folder{
	var folders []Folder
	for _, letter:=range pbLetter{
		letterModel:=Folder{Name: letter.Name, Type: letter.Type}
		folders=append(folders, letterModel)
	}
	return folders
}

func ProtoToModelMail(pbLetter *mailProto.LetterListResponse) []Model.Letter{
	var letters []Model.Letter
	for _, pb:=range pbLetter.Letter{
		letter:=Model.Letter{
			Sender: pb.Sender,
			Receiver: pb.Receiver,
			Text: pb.Text,
			Theme: pb.Theme,
			IsWatched: pb.IsWatched,
			Id: pb.Lid,
			DateTime: int64(pb.DateTime),
			DirectoryRecv: pb.DirectoryRecv,
		}
		letters=append(letters, letter)
	}
	return letters
}
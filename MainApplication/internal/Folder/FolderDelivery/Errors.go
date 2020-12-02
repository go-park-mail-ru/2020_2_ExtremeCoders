package FolderDelivery

import (
	mailProto "MainApplication/proto/MailService"
	userProto "MainApplication/proto/UserServise"
)

func ProtoFolderListResponse(folders []*userProto.FolderNameType) []byte {
	ans := FolderList{
		Code:    200,
		Folders: ProtoToModelList(folders),
	}
	res, err := ans.MarshalJSON()
	if err!=nil{
		return nil
	}
	return res
}

func SuccessRespAns() []byte {
	ans := SuccessAns{Code: 200}
	res, _ := ans.MarshalJSON()
	return res
}

func GetFoldersError(err error) []byte {
	ans := &LetterErr{
		Code:        400,
		Description: err.Error(),
	}
	res, _ := ans.MarshalJSON()
	return res
}

func ProtoResponseAnswer(pbLetter *mailProto.Response) []byte {
	code := 200
	if !pbLetter.Ok {
		code = 409
	}
	ans := LetterErr{
		Code:        code,
		Description: pbLetter.Description,
	}
	res, _ := ans.MarshalJSON()
	return res
}

func ProtoLetterListAnswer(pbLetter *mailProto.LetterListResponse) []byte {
	code := 200
	if !pbLetter.Result.Ok {
		code = 409
	}
	ans := LetterList{
		Code:        code,
		Description: pbLetter.Result.Description,
		letter:      pbLetter.Letter,
	}
	res, _ := ans.MarshalJSON()
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
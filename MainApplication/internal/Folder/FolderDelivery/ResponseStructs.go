package FolderDelivery

import (
	mailProto "MainApplication/proto/MailService"
	userProto "MainApplication/proto/UserServise"
)

//var getFolderListError = errors.New("getErrorListError")
type LetterErr struct {
	Code        int
	Description string
}

type LetterList struct {
	Code        int
	Description string
	letter      []*mailProto.Letter
}

type FolderList struct {
	Code    int
	Folders []Folder
}

type SuccessAns struct {
	Code int
}

type Folder struct {
	Name string
	Type string
}

func ProtoToModelList(pbLetter []*userProto.FolderNameType) []Folder{
	var folders []Folder
	for _, letter:=range pbLetter{
		letterModel:=Folder{Name: letter.Name, Type: letter.Type}
		folders=append(folders, letterModel)
	}
	return folders
}
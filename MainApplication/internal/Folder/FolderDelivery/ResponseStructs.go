package FolderDelivery

import (
	mailProto "MainApplication/proto/MailService"
)

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

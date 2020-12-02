package FolderDelivery

import (
	Model "MainApplication/internal/Letter/LetterModel"
)

type LetterErr struct {
	Code        int
	Description string
}

type LetterList struct {
	Code        int
	Description string
	Letter      []Model.Letter
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

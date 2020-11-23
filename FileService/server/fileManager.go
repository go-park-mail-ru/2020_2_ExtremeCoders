package main

import (
	fileProto "Mailer/FileService/proto"
	"fmt"
	"golang.org/x/net/context"
	"os"
)


type FileManager struct {
}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (sm *FileManager) SetAvatar(ctx context.Context,avatar *fileProto.Avatar) (*fileProto.Nothing, error) {
	fmt.Println("SEVER CALL SetAvatar")
	path := avatar.Email + avatar.FileName
	f, err := os.Create(path)
	if err != nil {
		return &fileProto.Nothing{Dummy: true}, err
	}
	defer f.Close()
	fmt.Println(avatar.Content)
	_ ,err = f.Write(avatar.Content)
	return &fileProto.Nothing{Dummy: true}, err
}

func (sm *FileManager) GetAvatar(context.Context, *fileProto.User) (*fileProto.Avatar, error) {
	fmt.Printf("SEVER CALL GetAvatar")
	return &fileProto.Avatar{Email: "s@mail.ru", FileName: "asdfa.png"}, nil
}

func (sm *FileManager) SaveFile(context.Context, *fileProto.File) (*fileProto.Nothing, error) {
	fmt.Printf("SEVER CALL SaveFile")
	return &fileProto.Nothing{Dummy: true}, nil
}

func (sm *FileManager) GetFile(context.Context, *fileProto.LetterId) (*fileProto.File, error) {
	fmt.Printf("SEVER CALL GetFile")
	return &fileProto.File{LetterId: 29, FileType: "png"}, nil
}

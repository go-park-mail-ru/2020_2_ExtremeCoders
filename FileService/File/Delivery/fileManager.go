package Delivery

import (
	"FileService/File/UseCase"
	fileProto "FileService/proto"
	"golang.org/x/net/context"
)

type FileManager struct {
	uc UseCase.Interface
}

func NewFileManager(uc UseCase.Interface) fileProto.FileServiceServer {
	return &FileManager{uc: uc}
}

func (fm *FileManager) SetAvatar(ctx context.Context, avatar *fileProto.Avatar) (*fileProto.Nothing, error) {
	err := fm.uc.SaveAvatar(avatar)
	return &fileProto.Nothing{Dummy: true}, err
}

func (fm *FileManager) GetAvatar(ctx context.Context, user *fileProto.User) (*fileProto.Avatar, error) {
	avatar, err := fm.uc.GetAvatar(user)
	return avatar, err
}

func (fm *FileManager) SaveFiles(ctx context.Context, files *fileProto.Files) (*fileProto.Nothing, error) {
	err := fm.uc.SaveFiles(files)
	return &fileProto.Nothing{Dummy: true}, err
}

func (fm *FileManager) GetFiles(ctx context.Context, id *fileProto.LetterId) (*fileProto.Files, error) {
	files, err := fm.uc.GetFiles(id)
	return files, err
}

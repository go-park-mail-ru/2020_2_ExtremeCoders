package UseCase

import fileProto "Mailer/FileService/proto"
import "Mailer/FileService/internal/File/Repository"
//go:generate mockgen -source=UseCase.go -destination=UseCaseMock.go

type Interface interface {
	SaveFiles(*fileProto.Files) error
	SaveAvatar(*fileProto.Avatar) error
	GetAvatar(user *fileProto.User) (*fileProto.Avatar, error)
	GetFiles(id *fileProto.LetterId) (*fileProto.Files, error)
}

type UseCase struct {
	repo Repository.Interface
}

func New(repo Repository.Interface) Interface {
	return UseCase{repo: repo}
}

func (uc UseCase) SaveFiles(file *fileProto.Files) error {
	return uc.repo.SaveFiles(file)
}

func (uc UseCase) SaveAvatar(avatar *fileProto.Avatar) error {
	return uc.repo.SaveAvatar(avatar)
}

func (uc UseCase) GetAvatar(user *fileProto.User) (*fileProto.Avatar, error) {
	avatar, err := uc.repo.GetAvatar(user)
	if err != nil {
		err = nil
		avatar, err = uc.repo.GetDefaultAvatar()
	}
	if err != nil {
		return &fileProto.Avatar{}, err
	}
	return avatar, err
}

func (uc UseCase) GetFiles(id *fileProto.LetterId) (*fileProto.Files, error) {
	return uc.repo.GetFiles(id)
}

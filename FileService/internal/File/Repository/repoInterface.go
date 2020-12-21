package Repository

import (
	fileProto "FileService/proto"
)
//go:generate mockgen -source=repoInterface.go -destination=RepositoryMock.go
type Interface interface {
	SaveFiles(*fileProto.Files) error
	GetFiles(*fileProto.LetterId) (*fileProto.Files, error)
	SaveAvatar(*fileProto.Avatar) error
	GetAvatar(*fileProto.User) (*fileProto.Avatar, error)
	GetDefaultAvatar() (*fileProto.Avatar, error)
}

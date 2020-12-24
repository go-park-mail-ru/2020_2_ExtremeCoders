package UserRepository

import (
	"Mailer/UserService/internal/UserModel"
	"errors"
)
//go:generate mockgen -destination=../../test/mock_Repository/mock_UserRepository.go -package=mocks -source=./DataBaseRequests.go

var EmailAlreadyExists = errors.New("Email already exist!")
var DbError = errors.New("Data Base error!")
var CantAddSession = errors.New("Add session error!")
var CantAddUser = errors.New("Add user error!")
var CantGetUserByEmail = errors.New("Get user by email error!")
var CantGetUserByUid = errors.New("Get user by uid error!")
var InvalidSession = errors.New("Invalid session!")
var CantGetUserOnUpdate = errors.New("Get user on update error!")
var CantUpdateUser = errors.New("User update error!")
var RemoveSessionError = errors.New("Could not remove session!")
var GetSessionError = errors.New("Could not get session!")
var GetFolderIdError = errors.New("Could not get folder ID!")
var CreateFolderError = errors.New("Could not get folder ID!")
var RenameFolderError = errors.New("Could not rename folder!")

type UserDB interface {
	IsEmailExists(email string) error
	AddSession(sid string, uid uint64, user *UserModel.User) error
	AddUser(user *UserModel.User) error
	GenerateSID() (sid []rune, err error)
	GenerateUID() (uid uint64, err error)
	GetUserByEmail(email string) (user *UserModel.User, err error)
	GetUserByUID(uid uint64) (user *UserModel.User, err error)
	IsOkSession(sid string) (uid uint64, err error)
	UpdateProfile(newUser UserModel.User, email string) error
	RemoveSession(sid string) (err error, uid uint64)
	GetSessionByUID(uid uint64) (sid string, err error)
	GetFolderId(uid uint64, kind string, name string) (fid uint64, err error)
	CreateFolder(name string, kind string, uid uint64) error
	RenameFolder(uid uint64, kind string, oldName string, newName string) error
	RemoveFolder(id uint64) error
	GetFoldersList(uid uint64, kind string) (folders []*UserModel.Folder, err error)
}

//type UserDB interface {
//	IsEmailExists(*proto.Email) error
//	AddSession(*proto.AddSessionMsg) error
//	AddUser(*proto.User) (*proto.Nothing, error)
//	GenerateSID(*proto.Nothing) (*proto.Sid, error)
//	GenerateUID(*proto.Nothing) (*proto.Sid, error)
//	GetUserByEmail(*proto.Email) (*proto.User, error)
//	GetUserByUID(*proto.Uid) (*proto.User, error)
//	IsOkSession(*proto.Sid) (*proto.Uid, error)
//	UpdateProfile(*proto.UpdateProfileMsg) (*proto.Nothing, error)
//	RemoveSession(*proto.Sid) (*proto.Uid, error)
//	GetSessionByUID(*proto.Uid) (*proto.Sid, error)
//}

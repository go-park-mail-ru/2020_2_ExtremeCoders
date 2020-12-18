package UserRepository

import (
	"Mailer/MainApplication/internal/User/UserModel"
	"errors"
)

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

//go:generate mockgen -source=./DataBaseRequests.go -destination=../../../test/mock_UserRepository/RepositoryMock.go

type UserDB interface {
	IsEmailExists(string) error
	AddSession(string, uint64, *UserModel.User) error
	AddUser(*UserModel.User) error
	GenerateSID() ([]rune, error)
	GenerateUID() (uint64, error)
	GetUserByEmail(string) (*UserModel.User, error)
	GetUserByUID(uint64) (*UserModel.User, error)
	IsOkSession(string) (uint64, error)
	UpdateProfile(UserModel.User, string) error
	RemoveSession(string) (error, uint64)
	GetSessionByUID(uint64) (string, error)
}

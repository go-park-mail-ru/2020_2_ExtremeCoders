package UserRepository

import (
	"CleanArch/internal/User/UserModel"
	"errors"
)

var EmailAlreadyExists=errors.New("Email already exist!")
var DbError=errors.New("Data Base error!")

var ReceiverNotFound=errors.New("Receiver not found!")
var SaveLetterError=errors.New("Save letter error!")


var GetUserError=errors.New("Could not get user!")
var RemoveSessionError =errors.New("Could not remove session!")
type UserDB interface {
	IsEmailExists(string) error
	AddSession(string, uint64, *UserModel.User) error
	AddUser(*UserModel.User, error)
	GenerateSID() ([]rune, error)
	GenerateUID() (uint64, error)
	GetUserByEmail(string) (*UserModel.User, error)
	GetUserByUID(uint64) (*UserModel.User, error)
	IsOkSession(string) (uint64,error)
	UpdateProfile(UserModel.User, string) error
	RemoveSession(uint64, string) error
}

package UserRepository

import "CleanArch/internal/User/UserModel"

type UserDB interface {
	IsEmailExists(string) bool
	AddSession(string, uint64, *UserModel.User) error
	AddUser(*UserModel.User)
	GenerateSID() []rune
	GenerateUID() uint64
	GetUserByEmail(string) (*UserModel.User, bool)
	GetUserByUID(uint64) *UserModel.User
	IsOkSession(string) (uint64,bool)
	UpdateProfile(UserModel.User, string)
	RemoveSession(uint64, string)
	RemoveSessionByUID(uint64)
	ShowAll()
}

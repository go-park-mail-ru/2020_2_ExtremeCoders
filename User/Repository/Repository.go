package Repository

import (
	"CleanArch/User/Models"
)

type DB interface {
	IsEmailExists(string) bool
	AddSession(string, uint64) error
	AddUser(Models.User)
	GenerateSID() []rune
	GenerateUID() uint64
	GetUserByEmail(string) (*Models.User, bool)
	GetUserByUID(uint64) *Models.User
	IsOkSession(string) (uint64,bool)
	UpdateProfile(Models.User, string)
	RemoveSession(uint64, string)
}

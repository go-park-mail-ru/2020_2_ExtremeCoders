package Repository

import (
	"CleanArch/app/Models"
)

type UserDB interface {
	IsEmailExists(string) bool
	AddSession(string, uint64, *Models.User) error
	AddUser(*Models.User)
	GenerateSID() []rune
	GenerateUID() uint64
	GetUserByEmail(string) (*Models.User, bool)
	GetUserByUID(uint64) *Models.User
	IsOkSession(string) (uint64,bool)
	UpdateProfile(Models.User, string)
	RemoveSession(uint64, string)
	RemoveSessionByUID(uint64)
	ShowAll()

	SaveMail(Models.Letter) int
	GetLetters(string) (int, []Models.Letter)
	GetSendedLetters(string) (int, []Models.Letter)
	GenerateLID() uint64
}

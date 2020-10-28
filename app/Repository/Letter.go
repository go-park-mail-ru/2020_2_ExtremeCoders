package Repository

import (
	Models2 "CleanArch/app/Models"
)

type LetterDB interface {
	IsEmailExists(string) bool
	SaveMail(letter Models2.Letter)
	GetMail(letter Models2.Letter)
}
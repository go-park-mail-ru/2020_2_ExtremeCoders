package LetterRepository

import (
	"CleanArch/cmd/Letter/LetterModel"
)

type LetterDB interface {
	SaveMail(LetterModel.Letter) int
	GetRecievedLetters(string) (int, []LetterModel.Letter)
	GetSendedLetters(string) (int, []LetterModel.Letter)
	GenerateLID() uint64
}

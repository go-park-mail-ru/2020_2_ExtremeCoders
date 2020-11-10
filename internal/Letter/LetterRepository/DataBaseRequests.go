package LetterRepository

import (
	"CleanArch/internal/Letter/LetterModel"
	"errors"
)


var DbError=errors.New("Data Base error!")
var ReceiverNotFound=errors.New("Receiver not found!")
var SaveLetterError=errors.New("Save letter error!")


type LetterDB interface {
	IsUserExist(email string) error
	SaveMail(LetterModel.Letter) error
	GetRecievedLetters(string) (error, []LetterModel.Letter)
	GetSendedLetters(string) (error, []LetterModel.Letter)
	GenerateLID() uint64
}

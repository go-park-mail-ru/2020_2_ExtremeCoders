package LetterRepository

import (
	"Mailer/MainApplication/internal/Letter/LetterModel"
	"errors"
)

var DbError = errors.New("Data Base error!")
var ReceiverNotFound = errors.New("Receiver not found!")
var SaveLetterError = errors.New("Save letter error!")

var ReceivedLetterError = errors.New("Could not get received letters!")
var SentLetterError = errors.New("Could not get sent letters!")

type LetterDB interface {
	IsUserExist(email string) error
	SaveMail(LetterModel.Letter) error
	GetReceivedLetters(string) (error, []LetterModel.Letter)
	GetSendedLetters(string) (error, []LetterModel.Letter)
	GenerateLID() uint64
}

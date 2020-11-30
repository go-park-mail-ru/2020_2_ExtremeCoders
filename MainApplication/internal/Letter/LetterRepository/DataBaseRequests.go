package LetterRepository

import (
	"MainApplication/internal/Letter/LetterModel"
	"errors"
)

var DbError = errors.New("Data Base error!")
var ReceiverNotFound = errors.New("Receiver not found!")
var SaveLetterError = errors.New("Save letter error!")

var ReceivedLetterError = errors.New("Could not get received letters!")
var SentLetterError = errors.New("Could not get sent letters!")
var WatchLetterError = errors.New("Could not watch letter!")

type LetterDB interface {
	SaveMail(LetterModel.Letter) error
	GetReceivedLetters(string) (error, []LetterModel.Letter)
	GetSendedLetters(string) (error, []LetterModel.Letter)

	GetReceivedLettersDir(uint64) (error, []LetterModel.Letter)
	GetSendedLettersDir(uint64) (error, []LetterModel.Letter)

	WatchLetter(uint64) (error, LetterModel.Letter)
}

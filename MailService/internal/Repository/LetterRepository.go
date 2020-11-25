package Repository

import (
	"MailService/internal/Model"
	"errors"
)

var DbError = errors.New("Data Base error!")
var ReceiverNotFound = errors.New("Receiver not found!")
var SaveLetterError = errors.New("Save letter error!")

var ReceivedLetterError = errors.New("Could not get received letters!")
var SentLetterError = errors.New("Could not get sent letters!")
var GetByLidError = errors.New("Could not get letter by lid!")
var SetLetterWatchedError = errors.New("Could not set letter watched!")


type LetterDB interface {
	SaveMail(Model.Letter) error
	GenerateLID() uint64
	SetLetterWatched(uint64) (error, Model.Letter)
	GetLetterByLid(uint64)(error, Model.Letter)

	GetLettersRecv(uint64)  (error, []Model.Letter)
	GetLettersSent(uint64)  (error, []Model.Letter)
}
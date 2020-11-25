package UseCase

import (
	"MailService/internal/Model"
	"MailService/internal/Repository"
)

type UseCase struct{
	re Repository.LetterDB
}

type Interface interface {
	GetLettersByDir(dir string) (error, []Model.Letter)
	SaveLetter(letter Model.Letter) error
	WatchLetter(lid uint64) (error, Model.Letter)
}

func (uc UseCase)GetLettersRecv(dir uint64) (error, []Model.Letter){
	err, letters:=uc.re.GetLettersRecv(dir)
	return err, letters
}

func (uc UseCase)GetLettersSend(dir uint64) (error, []Model.Letter){
	err, letters:=uc.re.GetLettersSent(dir)
	return err, letters
}

func (uc UseCase) SaveLetter(letter Model.Letter) error{
	err:=uc.re.SaveMail(letter)
	return err
}

func (uc UseCase) WatchLetter(lid uint64) (error, Model.Letter){
	return uc.re.SetLetterWatched(lid)
}
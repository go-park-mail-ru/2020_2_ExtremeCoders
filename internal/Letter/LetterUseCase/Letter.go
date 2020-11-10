package LetterUseCase

import (
	"CleanArch/internal/Letter/LetterModel"
	"CleanArch/internal/Letter/LetterRepository"
	"errors"
	"fmt"
)

type UseCase struct{
	Db LetterRepository.LetterDB
}

func (uc *UseCase) SaveLetter(letter *LetterModel.Letter) error {
	letter.Id = uc.Db.GenerateLID()
	err:=uc.Db.IsUserExist(letter.Receiver)
	if err!=nil{
		return err
	}
	err=uc.Db.SaveMail(*letter)
	if err!=nil{
		return err
	}
	return nil
}

func (uc *UseCase) GetRecievedLetters(email string) (int, []LetterModel.Letter) {
	return uc.Db.GetReceivedLetters(email)
}

func (uc *UseCase) GetSendedLetters(email string) (int, []LetterModel.Letter) {
	return uc.Db.GetSendedLetters(email)
}

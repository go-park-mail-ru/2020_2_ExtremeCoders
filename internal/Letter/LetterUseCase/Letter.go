package LetterUseCase

import (
	"CleanArch/internal/Letter/LetterModel"
	"CleanArch/internal/Letter/LetterRepository"
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

func (uc *UseCase) GetReceivedLetters(email string) (error, []LetterModel.Letter) {
	err, letters:= uc.Db.GetReceivedLetters(email)
	if err!=nil{
		return err, nil
	}
	return nil, letters
}

func (uc *UseCase) GetSendedLetters(email string) (error, []LetterModel.Letter) {
	err, letters:=uc.Db.GetSendedLetters(email)
	if err!=nil{
		return err, nil
	}
	return nil, letters
}

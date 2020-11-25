package LetterUseCase

import (
	"CleanArch/MainApplication/internal/Letter/LetterModel"
	"CleanArch/MainApplication/internal/Letter/LetterRepository"
)

type LetterUseCase interface {
	SaveLetter(letter *LetterModel.Letter) error
	GetReceivedLetters(email string) (error, []LetterModel.Letter)
	GetSendedLetters(email string) (error, []LetterModel.Letter)
}

type useCase struct {
	Db LetterRepository.LetterDB
}

func New(db LetterRepository.LetterDB) LetterUseCase {
	return useCase{Db: db}
}

func (uc useCase) SaveLetter(letter *LetterModel.Letter) error {
	letter.Id = uc.Db.GenerateLID()
	err := uc.Db.IsUserExist(letter.Receiver)
	if err != nil {
		return err
	}
	err = uc.Db.SaveMail(*letter)
	if err != nil {
		return err
	}
	return nil
}

func (uc useCase) GetReceivedLetters(email string) (error, []LetterModel.Letter) {
	err, letters := uc.Db.GetReceivedLetters(email)
	if err != nil {
		return err, nil
	}
	return nil, letters
}

func (uc useCase) GetSendedLetters(email string) (error, []LetterModel.Letter) {
	err, letters := uc.Db.GetSendedLetters(email)
	if err != nil {
		return err, nil
	}
	return nil, letters
}
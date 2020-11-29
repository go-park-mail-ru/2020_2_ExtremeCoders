package LetterUseCase

import (
	"MainApplication/internal/Letter/LetterModel"
	"MainApplication/internal/Letter/LetterRepository"
)

type LetterUseCase interface {
	SaveLetter(letter *LetterModel.Letter) error
	GetReceivedLetters(email string) (error, []LetterModel.Letter)
	GetSendedLetters(email string) (error, []LetterModel.Letter)
	GetReceivedLettersDir(dir uint64) (error, []LetterModel.Letter)
	GetSendedLettersDir(dir uint64) (error, []LetterModel.Letter)
	WatchLetter(lid uint64)(error, LetterModel.Letter)
}

type useCase struct {
	Db LetterRepository.LetterDB
}

func New(db LetterRepository.LetterDB) LetterUseCase {
	return useCase{Db: db}
}

func (uc useCase) SaveLetter(letter *LetterModel.Letter) error {
	err := uc.Db.SaveMail(*letter)
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

func (uc useCase)WatchLetter(lid uint64)(error, LetterModel.Letter){
	err, letters := uc.Db.WatchLetter(lid)
	if err != nil {
		return err, LetterModel.Letter{}
	}
	return nil, letters
}

func (uc useCase)GetReceivedLettersDir(dir uint64) (error, []LetterModel.Letter){
	err, letters := uc.Db.GetReceivedLettersDir(dir)
	if err != nil {
		return err, nil
	}
	return nil, letters
}

func (uc useCase)GetSendedLettersDir(dir uint64) (error, []LetterModel.Letter){
	err, letters := uc.Db.GetSendedLettersDir(dir)
	if err != nil {
		return err, nil
	}
	return nil, letters
}
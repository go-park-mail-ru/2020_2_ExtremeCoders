package LetterUseCase

import (
	"Mailer/MainApplication/internal/Letter/LetterModel"
	"Mailer/MainApplication/internal/Letter/LetterRepository"
)
//go:generate mockgen -source=./Letter.go -destination=../../../test/mock_LetterUseCase/LetterUseCaseMock.go

type LetterUseCase interface {
	SaveLetter(letter *LetterModel.Letter) error
	GetReceivedLetters(email string, limit uint64, offset uint64) (error, []LetterModel.Letter)
	GetSendedLetters(email string, limit uint64, offset uint64) (error, []LetterModel.Letter)
	GetReceivedLettersDir(dir uint64, limit uint64, offset uint64) (error, []LetterModel.Letter)
	GetSendedLettersDir(dir uint64) (error, []LetterModel.Letter)
	WatchLetter(lid uint64) (error, LetterModel.Letter)
	DeleteLetter(lid uint64) error
	FindSim(sim string, email string) string
	GetLetterBy(what string, val string, email string, limit int, offset int) (error, []LetterModel.Letter)
	SetLetterInSpam(lid uint64) error
	SetLetterInBox(lid uint64) error
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

func (uc useCase) GetReceivedLetters(email string, limit uint64, offset uint64) (error, []LetterModel.Letter) {
	err, letters := uc.Db.GetReceivedLetters(email, limit, offset)
	if err != nil {
		return err, nil
	}
	return nil, letters
}

func (uc useCase) GetSendedLetters(email string, limit uint64, offset uint64) (error, []LetterModel.Letter) {
	err, letters := uc.Db.GetSendedLetters(email, limit, offset)
	if err != nil {
		return err, nil
	}
	return nil, letters
}

func (uc useCase) WatchLetter(lid uint64) (error, LetterModel.Letter) {
	err, letters := uc.Db.WatchLetter(lid)
	if err != nil {
		return err, LetterModel.Letter{}
	}
	return nil, letters
}

func (uc useCase) GetReceivedLettersDir(dir uint64, limit uint64, offset uint64) (error, []LetterModel.Letter) {
	err, letters := uc.Db.GetReceivedLettersDir(dir, limit, offset)
	if err != nil {
		return err, nil
	}
	return nil, letters
}

func (uc useCase) GetSendedLettersDir(dir uint64) (error, []LetterModel.Letter) {
	err, letters := uc.Db.GetSendedLettersDir(dir)
	if err != nil {
		return err, nil
	}
	return nil, letters
}

func (uc useCase) DeleteLetter(lid uint64) error{
	return uc.Db.DeleteLetter(lid)
}

func (uc useCase) FindSim(sim string, email string) string{
	return uc.Db.FindSimilar(sim, email)
}

func (uc useCase) GetLetterBy(what string, val string, email string, limit int, offset int) (error, []LetterModel.Letter){
	return uc.Db.GetLetterBy(what, val, email, limit, offset)
}

func (uc useCase)SetLetterInSpam(lid uint64) error{
	return uc.Db.SetLetterInSpam(lid)
}

func (uc useCase)SetLetterInBox(lid uint64) error{
	return uc.Db.SetLetterInBox(lid)
}

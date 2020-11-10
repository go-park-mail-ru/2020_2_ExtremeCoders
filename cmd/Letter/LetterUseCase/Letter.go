package LetterUseCase

import (
	"CleanArch/cmd/Letter/LetterModel"
	"CleanArch/cmd/Letter/LetterRepository"
)

type UseCase struct{
	Db LetterRepository.LetterDB
}

func (uc *UseCase) SaveLetter(letter *LetterModel.Letter) int {
	letter.Id = uc.Db.GenerateLID()
	return uc.Db.SaveMail(*letter)
}

func (uc *UseCase) GetRecievedLetters(email string) (int, []LetterModel.Letter) {
	return uc.Db.GetRecievedLetters(email)
}

func (uc *UseCase) GetSendedLetters(email string) (int, []LetterModel.Letter) {
	return uc.Db.GetSendedLetters(email)
}

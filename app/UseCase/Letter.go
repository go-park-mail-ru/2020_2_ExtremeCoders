package UseCase

import "CleanArch/app/Models"

func (uc *UseCase) SaveLetter(letter *Models.Letter) int {
	letter.Id = uc.Db.GenerateLID()
	return uc.Db.SaveMail(*letter)
}

func (uc *UseCase) GetRecievedLetters(email string) (int, []Models.Letter) {
	return uc.Db.GetRecievedLetters(email)
}

func (uc *UseCase) GetSendedLetters(email string) (int, []Models.Letter) {
	return uc.Db.GetSendedLetters(email)
}

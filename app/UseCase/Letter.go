package UseCase

import "CleanArch/app/Models"

func (uc *UseCase)SaveLetter(letter Models.Letter) int {
	letter.Id=uc.Db.GenerateLID()
	return uc.Db.SaveMail(letter)
}

func (uc *UseCase)GetLetters(email string)(int, []Models.Letter) {
	return uc.Db.GetLetters(email)
}
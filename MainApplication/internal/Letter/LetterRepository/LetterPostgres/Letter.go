package LetterPostgres
//
//import (
//	"MainApplication/internal/Letter/LetterModel"
//	"MainApplication/internal/Letter/LetterRepository"
//	pgwrapper "gitlab.com/slax0rr/go-pg-wrapper"
//)
//
//type dataBase struct {
//	DB pgwrapper.DB
//}
//
//func New(db pgwrapper.DB) LetterRepository.LetterDB {
//	return dataBase{DB: db}
//}
//
//func (dbInfo dataBase) SaveMail(letter LetterModel.Letter) error {
//	_, err := dbInfo.DB.Model(&letter).Insert()
//	if err != nil {
//		return LetterRepository.SaveLetterError
//	}
//	return nil
//}
//
//func (dbInfo dataBase) GetReceivedLetters(email uint64) (error, []LetterModel.Letter) {
//	var letters []LetterModel.Letter
//	exist := dbInfo.DB.Model(&letters).Where("receiver=?", email).Select()
//	if exist != nil {
//		return LetterRepository.ReceivedLetterError, nil
//	}
//	return nil, letters
//}
//
//func (dbInfo dataBase) GetSendedLetters(email uint64) (error, []LetterModel.Letter) {
//	var letters []LetterModel.Letter
//	exist := dbInfo.DB.Model(&letters).Where("sender=?", email).Select()
//	if exist != nil {
//		return LetterRepository.SentLetterError, nil
//	}
//	return nil, letters
//}
//
//func (dbInfo dataBase) WatchLetter(uint64) (error, LetterModel.Letter) {
//	return nil, LetterModel.Letter{}
//}

package LetterPostgres

import (
	"Mailer/MainApplication/internal/Letter/LetterModel"
	"Mailer/MainApplication/internal/Letter/LetterRepository"
	"Mailer/MainApplication/internal/User/UserModel"
	crypto "crypto/rand"
	pgwrapper "gitlab.com/slax0rr/go-pg-wrapper"
	"math/big"
)

type dataBase struct {
	DB pgwrapper.DB
}

func New(db pgwrapper.DB) LetterRepository.LetterDB {
	return dataBase{DB: db}
}

func (dbInfo dataBase) SaveMail(letter LetterModel.Letter) error {
	_, err := dbInfo.DB.Model(&letter).Insert()
	if err != nil {
		return LetterRepository.SaveLetterError
	}
	return nil
}

func (dbInfo dataBase) IsUserExist(email string) error {
	reciever := &UserModel.User{Email: email}
	erro := dbInfo.DB.Model(reciever).Where("email=?", email).Select() //uc
	if erro != nil {
		return LetterRepository.ReceiverNotFound
	}
	return nil
}

func (dbInfo dataBase) GetReceivedLetters(email string) (error, []LetterModel.Letter) {
	var letters []LetterModel.Letter
	exist := dbInfo.DB.Model(&letters).Where("receiver=?", email).Select()
	if exist != nil {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, letters
}

func (dbInfo dataBase) GenerateLID() uint64 {
	for {
		lid, _ := crypto.Int(crypto.Reader, big.NewInt(4294967295))
		user := LetterModel.Letter{Id: lid.Uint64()}
		exist := dbInfo.DB.Model(user).Where("id=?", lid.Int64()).Select()
		if exist != nil {
			return lid.Uint64()
		}
	}
}

func (dbInfo dataBase) GetSendedLetters(email string) (error, []LetterModel.Letter) {
	var letters []LetterModel.Letter
	exist := dbInfo.DB.Model(&letters).Where("sender=?", email).Select()
	if exist != nil {
		return LetterRepository.SentLetterError, nil
	}
	return nil, letters
}

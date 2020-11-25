package LetterPostgres

import (
	"MailService/internal/Model"
	"MailService/internal/Repository"
	crypto "crypto/rand"
	pgwrapper "gitlab.com/slax0rr/go-pg-wrapper"
	"math/big"
)

type dataBase struct {
	DB pgwrapper.DB
}

func New(db pgwrapper.DB) Repository.LetterDB {
	return dataBase{DB: db}
}

func (dbInfo dataBase) SaveMail(letter Model.Letter) error {
	_, err := dbInfo.DB.Model(&letter).Insert()
	if err != nil {
		return Repository.SaveLetterError
	}
	return nil
}

func (dbInfo dataBase) GenerateLID() uint64 {
	for {
		lid, _ := crypto.Int(crypto.Reader, big.NewInt(4294967295))
		user := Model.Letter{Id: lid.Uint64()}
		exist := dbInfo.DB.Model(user).Where("id=?", lid.Int64()).Select()
		if exist != nil {
			return lid.Uint64()
		}
	}
}

func (dbInfo dataBase) GetLettersByFolder(email string) (error, []Model.Letter) {
	var letters []Model.Letter
	exist := dbInfo.DB.Model(&letters).Where("sender=?", email).Select()
	if exist != nil {
		return Repository.SentLetterError, nil
	}
	return nil, letters
}

func (dbInfo dataBase)SetLetterWatched(lid uint64) (error, Model.Letter){
	err, letter:=dbInfo.GetLetterByLid(lid)
	if err!=nil{
		return Repository.GetByLidError, letter
	}
	letter.IsWatched=true
	_, err =dbInfo.DB.Model(letter).Column("iswatched").Where("id=?", lid).Update()
	if err!=nil{
		return Repository.SetLetterWatchedError, letter
	}
	return nil, letter
}

func (dbInfo dataBase)GetLetterByLid(lid uint64)(error, Model.Letter){
	var letter Model.Letter
	exist := dbInfo.DB.Model(&letter).Where("id=?", lid).Select()
	if exist != nil {
		return Repository.SentLetterError, letter
	}
	return nil, letter
}

func (dbInfo dataBase)GetLettersRecv(Did uint64)  (error, []Model.Letter){
	var letters []Model.Letter
	exist := dbInfo.DB.Model(&letters).Where("directoryrecv=?", Did).Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}
	return nil, letters
}

func (dbInfo dataBase)GetLettersSent(Did uint64)  (error, []Model.Letter){
	var letters []Model.Letter
	exist := dbInfo.DB.Model(&letters).Where("directorysend=?", Did).Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}
	return nil, letters
}

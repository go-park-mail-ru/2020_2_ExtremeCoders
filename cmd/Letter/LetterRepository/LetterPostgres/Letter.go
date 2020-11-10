package LetterPostgres

import (
	"CleanArch/cmd/Letter/LetterModel"
	"CleanArch/cmd/Postgres"
	"CleanArch/cmd/User/UserModel"
	crypto "crypto/rand"
	"math/big"
)

func (dbInfo Postgres.DataBase)SaveMail(letter LetterModel.Letter)int {
	reciever:=&UserModel.User{Email: letter.Receiver}
	erro:=dbInfo.DB.Model(reciever).Where("email=?", letter.Receiver).Select()
	if erro!=nil{return 408}
	_, err:=dbInfo.db.Model(&letter).Insert()
	if err!=nil{
		return 409
	}
	return 200
}

func (dbInfo Postgres.DataBase) GetRecievedLetters(email string) (int, []LetterModel.Letter){
	var letters []LetterModel.Letter
	exist := dbInfo.db.Model(&letters).Where("receiver=?", &email).Select()
	if exist!=nil{
		return 409, nil
	}
	return 200, letters
}

func (dbInfo Postgres.DataBase) GenerateLID() uint64 {
	for {
		lid,_ :=crypto.Int(crypto.Reader, big.NewInt(4294967295))
		user := LetterModel.Letter{Id: lid.Uint64()}
		exist := dbInfo.db.Model(user).Where("id=?", lid.Int64()).Select()
		if exist != nil {
			return lid.Uint64()
		}
	}
}


func (dbInfo Postgres.DataBase) GetSendedLetters(email string) (int, []LetterModel.Letter) {
	var letters []LetterModel.Letter
	exist := dbInfo.db.Model(&letters).Where("sender=?", &email).Select()
	if exist!=nil{
		return 409, nil
	}
	return 200, letters
}
package Postgres

import (
	Models "CleanArch/app/Models"
	crypto "crypto/rand"
	"math/big"
)

func (dbInfo DataBase)SaveMail(letter Models.Letter)int {
	reciever:=&Models.User{Email: letter.Receiver}
	erro:=dbInfo.db.Model(reciever).Where("email=?", letter.Receiver).Select()
	if erro!=nil{return 408}
	_, err:=dbInfo.db.Model(&letter).Insert()
	if err!=nil{
		return 409
	}
	return 200
}

func (dbInfo DataBase)GetLetters(email string) (int, []Models.Letter){
	var letters []Models.Letter
	exist := dbInfo.db.Model(&letters).Where("sender=?", &email).Select()
	if exist!=nil{
		return 409, nil
	}
	return 200, letters
}

func (dbInfo DataBase) GenerateLID() uint64 {
	for {
		lid,_ :=crypto.Int(crypto.Reader, big.NewInt(4294967295))
		user := Models.Letter{Id: lid.Uint64()}
		exist := dbInfo.db.Model(user).Where("id=?", lid.Int64()).Select()
		if exist != nil {
			return lid.Uint64()
		}
	}
}


func (dbInfo DataBase) GetSendedLetters(email string) (int, []Models.Letter) {
	var letters []Models.Letter
	exist := dbInfo.db.Model(&letters).Where("receiver=?", &email).Select()
	if exist!=nil{
		return 409, nil
	}
	return 200, letters
}
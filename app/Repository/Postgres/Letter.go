package Postgres

import (
	Models "CleanArch/app/Models"
	crypto "crypto/rand"
	"math/big"
)

func (dbInfo DataBase)SaveMail(letter Models.Letter)int {
	_, err:=dbInfo.db.Model(&letter).Insert()
	if err!=nil{
		return 409
	}
	return 200
}

func (dbInfo DataBase)GetLetters(email string) (int, []Models.Letter){
	var letters []Models.Letter
	exist := dbInfo.db.Model(letters).Where("email=?", email).Select()
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
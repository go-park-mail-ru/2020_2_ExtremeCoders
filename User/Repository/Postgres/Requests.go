package Postgres

import (
	"CleanArch/User/Models"
	"math/rand"
)

const (
	SizeSID = 32
)

var SidRunes = []rune("1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

func (dbInfo DataBase)IsEmailExists(email string) bool{
	user :=&Models.User{Email: email}
	err:=dbInfo.db.Model(user).Select()
	if err!=nil{
		return false
	}
	return true
}

func (dbInfo DataBase)AddUser(user Models.User){
	_, err := dbInfo.db.Model(user).Insert()
	if err != nil {
		panic(err)
	}
}

func (dbInfo DataBase)GenerateSID() []rune{
	var sid = make([]rune, SizeSID)
	for {
		for i := 0; i < SizeSID; i++ {
			sid[i] = SidRunes[rand.Intn(len(SidRunes))]
		}
		session:= &Models.Session{Sid: string(sid)}
		exist:=dbInfo.db.Model(session).Select()
		if exist==nil {
			break
		}
	}
	return sid
}

func (dbInfo DataBase)GenerateUID() uint64 {
	var uid uint64
	for {
		for i := 0; i < SizeSID; i++ {
			uid = rand.Uint64()
		}
		user:=Models.User{Id: uid}
		exist:=dbInfo.db.Model(user).Select()
		if exist==nil {
			break
		}
	}
	return uid
}

func (dbInfo DataBase)GetUserByEmail(email string) (*Models.User, bool){
	user :=&Models.User{Email: email}
	err:=dbInfo.db.Model(user).Select()
	if err==nil{
		return user, true
	}
	return user, false
}

func (dbInfo DataBase)GetUserByUID(uid uint64) *Models.User{
	user :=&Models.User{Id: uid}
	dbInfo.db.Model(user).Select()
	return user
}

func (dbInfo DataBase)IsOkSession(sid string) (uint64,bool){
	session :=&Models.Session{Sid: sid}
	err:=dbInfo.db.Model(session).Select()
	if err!=nil{
		return 0, false
	}
	return session.User.Id, true
}

func (dbInfo DataBase)UpdateProfile(newUser Models.User, email string){
	oldUser:=Models.User{Email: email}
	dbInfo.db.Model(oldUser).Select()

	User:= oldUser
	User.Name=newUser.Name
	User.Surname=newUser.Surname
	_, _ = dbInfo.db.Model(User).Column("Name", "Surname").Update()

}

func (dbInfo DataBase)RemoveSession(uid uint64, sid string){
	session:=&Models.Session{Sid: sid}
	dbInfo.db.Model(session).WherePK().Delete()
}
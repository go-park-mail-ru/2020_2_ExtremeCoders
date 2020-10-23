package Postgres

import (
	"CleanArch/User/Models"
	"fmt"
	"math/rand"
)

const (
	SizeSID = 32
)

var SidRunes ="1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func (dbInfo DataBase) IsEmailExists(email string) bool {
	user := &Models.User{Email: email}
	err := dbInfo.db.Model(user).Where("email=?", email).Select()
	if err !=nil {
		return false
	}
	return true
}

func (dbInfo DataBase) AddUser(user *Models.User) {
	_, err := dbInfo.db.Model(user).Insert()
	if err != nil {

	}
}

func (dbInfo DataBase) AddSession(sid string, uid uint64, user *Models.User) error {
	session := &Models.Session{Id: sid, UserId: int64(uid), User: user}
	_, err := dbInfo.db.Model(session).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (dbInfo DataBase) GenerateSID() []rune {
	var sid string
	for {
		for i := 0; i < SizeSID; i++ {
			sid+= string(SidRunes[rand.Intn(len(SidRunes))])
		}
		fmt.Println(sid)
		session := &Models.Session{Id: sid}
		exist := dbInfo.db.Model(session).WherePK().Select()
		if exist != nil {
			break
		}
		sid=""
	}
	return []rune(sid)
}

func (dbInfo DataBase) GenerateUID() uint64 {
	var uid uint64
	for {
		uid = rand.Uint64()
		user := Models.User{Id: uid}
		exist := dbInfo.db.Model(user).Select()
		if exist != nil {
			break
		}
	}
	return uid
}

func (dbInfo DataBase) GetUserByEmail(email string) (*Models.User, bool) {
	user := &Models.User{Email: email}
	err := dbInfo.db.Model(user).Where("email=?", email).Select()
	if err == nil {
		return user, true
	}
	return user, false
}

func (dbInfo DataBase) GetUserByUID(uid uint64) *Models.User {
	user := &Models.User{Id: uid}
	dbInfo.db.Model(user).Select()
	return user
}

func (dbInfo DataBase) IsOkSession(sid string) (uint64, bool) {
	session := &Models.Session{Id: sid}
	err := dbInfo.db.Model(session).WherePK().Select()
	if err != nil  {
		return 0, false
	}
	return uint64(session.UserId), true
}

func (dbInfo DataBase) UpdateProfile(newUser Models.User, email string) {
	oldUser := Models.User{Email: email}
	dbInfo.db.Model(oldUser).Select()

	User := oldUser
	User.Name = newUser.Name
	User.Surname = newUser.Surname
	_, _ = dbInfo.db.Model(User).Column("Name", "Surname").Update()

}

func (dbInfo DataBase) RemoveSession(uid uint64, sid string) {
	session := &Models.Session{Id: sid}
	dbInfo.db.Model(session).WherePK().Delete()
}

func (dbInfo DataBase) RemoveSessionByUID(uid uint64){
	session:=&Models.Session{UserId: int64(uid)}
	err:=dbInfo.db.Model(session).Where("user_id=?", uid).Select()
	if err!=nil{
		return
	}
	dbInfo.db.Model(session).WherePK().Delete()
}

func (dbInfo DataBase) ShowAll(){
	var users []Models.User
	var sessions []Models.Session
	fmt.Println(dbInfo.db.Model(users).Select())
	fmt.Println(dbInfo.db.Model(sessions).Select())
}
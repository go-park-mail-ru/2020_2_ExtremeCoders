package UserPostgres

import (
	"CleanArch/internal/User/UserModel"
	"CleanArch/internal/User/UserRepository"
	crypto "crypto/rand"
	"fmt"
	"github.com/go-pg/pg/v10"
	"math/big"
)

type DataBase struct {
	DB *pg.DB
}

const (
	SizeSID = 32
)

var SidRunes = "1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func (dbInfo DataBase) IsEmailExists(email string) error {
	user := &UserModel.User{Email: email}
	err := dbInfo.DB.Model(user).Where("email=?", email).Select()
	if err != nil {
		return UserRepository.EmailAlreadyExists
	}
	return nil
}

func (dbInfo DataBase) AddUser(user *UserModel.User) error {
	_, err := dbInfo.DB.Model(user).Insert()
	if err != nil {
		return UserRepository.CantAddUser
	}
	return nil
}

func (dbInfo DataBase) AddSession(sid string, uid uint64, user *UserModel.User) error {
	session := &UserModel.Session{Id: sid, UserId: int64(uid), User: user}
	_, err := dbInfo.DB.Model(session).Insert()
	if err != nil {
		return UserRepository.CantAddSession
	}
	return nil
}

func (dbInfo DataBase) GenerateSID() ([]rune, error) {
	var sid string
	for {
		for i := 0; i < SizeSID; i++ {
			safeNum, _ := crypto.Int(crypto.Reader, big.NewInt(int64(len(SidRunes))))
			sid += string(SidRunes[safeNum.Int64()])
		}
		fmt.Println(sid)
		session := &UserModel.Session{Id: sid}
		exist := dbInfo.DB.Model(session).WherePK().Select()
		if exist != nil {
			break
		}
		sid = ""
	}
	return []rune(sid), nil
}

func (dbInfo DataBase) GenerateUID() (uint64, error) {
	for {
		uid, _ := crypto.Int(crypto.Reader, big.NewInt(4294967295))
		user := UserModel.User{Id: uid.Uint64()}
		exist := dbInfo.DB.Model(user).Where("id=?", uid.Int64()).Select()
		if exist != nil {
			return uid.Uint64(), nil
		}
	}
}

func (dbInfo DataBase) GetUserByEmail(email string) (*UserModel.User, error) {
	user := &UserModel.User{Email: email}
	err := dbInfo.DB.Model(user).Where("email=?", email).Select()
	if err == nil {
		return user, UserRepository.CantGetUserByEmail
	}
	return user, nil
}

func (dbInfo DataBase) GetUserByUID(uid uint64) (*UserModel.User, error) {
	user := &UserModel.User{Id: uid}
	err := dbInfo.DB.Model(user).WherePK().Select()
	if err == nil {
		return user, UserRepository.CantGetUserByUid
	}
	return user, nil
}

func (dbInfo DataBase) IsOkSession(sid string) (uint64, error) {
	session := &UserModel.Session{Id: sid}
	err := dbInfo.DB.Model(session).WherePK().Select()
	if err != nil {
		return 0, UserRepository.InvalidSession
	}
	return uint64(session.UserId), nil
}

func (dbInfo DataBase) UpdateProfile(newUser UserModel.User, email string) error {
	oldUser := &UserModel.User{Email: email}
	err := dbInfo.DB.Model(oldUser).Where("email=?", email).Select()
	if err != nil{
		return UserRepository.CantGetUserOnUpdate
	}
	User := oldUser
	User.Name = newUser.Name
	User.Surname = newUser.Surname
	User.Img = newUser.Img
	_, err = dbInfo.DB.Model(User).Column("name", "surname", "img").Where("email=?", email).Update()
	if err != nil{
		return UserRepository.CantUpdateUser
	}
	return nil
}

func (dbInfo DataBase) RemoveSession(uid uint64, sid string) error {
	session := &UserModel.Session{Id: sid}
	_, err := dbInfo.DB.Model(session).WherePK().Delete()
	if err != nil {
		return UserRepository.RemoveSessionError
	}
	return nil
}


func (dbInfo DataBase) GetSessionByUID(uid uint64) (string,error){
	session:=&UserModel.Session{UserId: int64(uid)}
	err:=dbInfo.DB.Model(session).Where("user_id=?", uid).Select()
	if err!=nil{
		return "",UserRepository.GetSessionError
	}
	return session.Id,nil
}


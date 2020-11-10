package UserPostgres

import (
	"CleanArch/cmd/Postgres"
	"CleanArch/cmd/User/UserModel"
)

const (
	SizeSID = 32
)

var SidRunes = "1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"


func (dbInfo Postgres.DataBase) IsEmailExists(email string) bool {
	user := &UserModel.User{Email: email}
	err := dbInfo.DB.Model(user).Where("email=?", email).Select()
	if err !=nil {
		return false
	}
	return true
}

//func (dbInfo DataBase) AddUser(user *UserModel.User) {
//	_, err := dbInfo.db.Model(user).Insert()
//	if err != nil {
//
//	}
//}
//
//func (dbInfo DataBase) AddSession(sid string, uid uint64, user *UserModel.User) error {
//	session := &UserModel.Session{Id: sid, UserId: int64(uid), User: user}
//	_, err := dbInfo.db.Model(session).Insert()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (dbInfo DataBase) GenerateSID() []rune {
//	var sid string
//	for {
//		for i := 0; i < SizeSID; i++ {
//			safeNum, _ := crypto.Int(crypto.Reader, big.NewInt(int64(len(SidRunes))))
//			sid += string(SidRunes[safeNum.Int64()])
//		}
//		fmt.Println(sid)
//		session := &UserModel.Session{Id: sid}
//		exist := dbInfo.db.Model(session).WherePK().Select()
//		if exist != nil {
//			break
//		}
//		sid = ""
//	}
//	return []rune(sid)
//}
//
//func (dbInfo DataBase) GenerateUID() uint64 {
//	for {
//		uid,_ :=crypto.Int(crypto.Reader, big.NewInt(4294967295))
//		user := UserModel.User{Id: uid.Uint64()}
//		exist := dbInfo.db.Model(user).Where("id=?", uid.Int64()).Select()
//		if exist != nil {
//			return uid.Uint64()
//		}
//	}
//}
//
//func (dbInfo DataBase) GetUserByEmail(email string) (*UserModel.User, bool) {
//	user := &UserModel.User{Email: email}
//	err := dbInfo.db.Model(user).Where("email=?", email).Select()
//	if err == nil {
//		return user, true
//	}
//	return user, false
//}
//
//func (dbInfo DataBase) GetUserByUID(uid uint64) *UserModel.User {
//	user := &UserModel.User{Id: uid}
//	dbInfo.db.Model(user).WherePK().Select()
//	return user
//}
//
//func (dbInfo DataBase) IsOkSession(sid string) (uint64, bool) {
//	session := &UserModel.Session{Id: sid}
//	err := dbInfo.db.Model(session).WherePK().Select()
//	if err != nil {
//		return 0, false
//	}
//	return uint64(session.UserId), true
//}
//
//func (dbInfo DataBase) UpdateProfile(newUser UserModel.User, email string) {
//	oldUser := &UserModel.User{Email: email}
//	dbInfo.db.Model(oldUser).Where("email=?", email).Select()
//
//	User := oldUser
//	User.Name = newUser.Name
//	User.Surname = newUser.Surname
//	User.Img = newUser.Img
//	_, err := dbInfo.db.Model(User).Column("name", "surname", "img").Where("email=?", email).Update()
//	fmt.Println(err)
//}
//
//func (dbInfo DataBase) RemoveSession(uid uint64, sid string) {
//	session := &UserModel.Session{Id: sid}
//	dbInfo.db.Model(session).WherePK().Delete()
//}
//
//func (dbInfo DataBase) RemoveSessionByUID(uid uint64){
//	session:=&UserModel.Session{UserId: int64(uid)}
//	err:=dbInfo.db.Model(session).Where("user_id=?", uid).Select()
//	if err!=nil{
//		return
//	}
//	dbInfo.db.Model(session).WherePK().Delete()
//}
//
//func (dbInfo DataBase) ShowAll(){
//	var users []UserModel.User
//	var sessions []UserModel.Session
//	fmt.Println(dbInfo.db.Model(users).Select())
//	fmt.Println(dbInfo.db.Model(sessions).Select())
//}

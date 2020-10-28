package tmpDB

import (
	"CleanArch/app/Models"
	"math/rand"
	"strconv"
)

const (
	SizeSID = 32
)

var SidRunes = []rune("1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

type LoggedIn struct {
	Sessions map[string]uint64
	Users    map[string] *Models.User
}

func (db LoggedIn)IsEmailExists(email string) bool{
	_, erro := db.Users[email]
	if erro {
		return true
	}
	return false
}

func (db LoggedIn)AddSession(sid string, uid uint64) error {
	db.Sessions[sid] = uid
	return  nil
}

func (db LoggedIn)AddUser(user Models.User){
	db.Users[user.Email] = &user
}

func (db LoggedIn)GenerateSID() []rune {
	var sid = make([]rune, SizeSID)
	for {
		for i := 0; i < SizeSID; i++ {
			sid[i] = SidRunes[rand.Intn(len(SidRunes))]
		}
		_, exist := db.Sessions[string(sid)]
		if !exist {
			break
		}
	}
	return sid
}

func (db LoggedIn)GenerateUID() uint64 {
	var uid uint64
	for {
		for i := 0; i < SizeSID; i++ {
			uid = rand.Uint64()
		}
		var _, exist = db.Users[strconv.FormatUint(uid, 10)]
		if !exist {
			break
		}
	}
	return uid
}

func (db LoggedIn)GetUserByEmail(email string) (*Models.User, bool){
	user, err:= db.Users[email]
	return user, err
}

func (db LoggedIn)GetUserByUID(uid uint64) *Models.User{
	for _, val := range db.Users {
		if (*val).Id == uid {
			return val
		}
	}
	return nil
}

func (db LoggedIn)IsOkSession(sid string) (uint64,bool){
	uid, ok := db.Sessions[sid]
	return uid, ok
}

func (db LoggedIn)UpdateProfile(newUser Models.User, email string){
	User, _:= db.Users[email]
	User.Name=newUser.Name
	User.Surname=newUser.Surname
	db.Users[email]=User
}

func (db LoggedIn)RemoveSession(uid uint64, sid string){
	User:=db.GetUserByUID(uid)
	delete(db.Sessions, sid)
	delete(db.Users, User.Email)
}
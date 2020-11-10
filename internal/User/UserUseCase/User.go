package UserUseCase

import (
	"CleanArch/internal/User/UserModel"
	"CleanArch/internal/User/UserRepository"
	"CleanArch/internal/errors"
	err "errors"
	"net/http"
	"time"
)

type UseCase struct{
	Db UserRepository.UserDB
}

var WrongPasswordError = err.New("Wrong password!")

func (uc *UseCase)Signup(user UserModel.User) (error, *http.Cookie) {
	if uc.Db.IsEmailExists(user.Email){
		//return 401, nil
		return UserRepository.EmailAlreadyExists, nil
	}

	user.Id = uc.Db.GenerateUID()
	sid := string(uc.Db.GenerateSID())
	uc.Db.AddUser(&user)
	if uc.Db.AddSession(sid, user.Id, &user) != nil {
		//return 401, nil
		return UserRepository.DbError, nil
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(24 * 7 * 4 * time.Hour),
	}
	cookie.Path = "/"

	//return 200, cookie
	return nil, cookie
}

func (uc *UseCase)SignIn(user UserModel.User) (error, string) {
	userEx, erro := uc.Db.GetUserByEmail(user.Email)
	if erro!=nil {
		return erro,""
	}
	if userEx.Password != user.Password {
		return WrongPasswordError, ""
	}
	sid,e := uc.Db.GenerateSID()
	if e!=nil{
		return e, ""
	}
	oldSid, er:=uc.Db.GetSessionByUID(userEx.Id)
	if er!=nil{
		return er, ""
	}
	uc.Db.RemoveSession(userEx.Id, oldSid)
	er=uc.Db.AddSession(string(sid), userEx.Id, &user)
	if er != nil {
		return er, ""
	}
	return nil, string(sid)

}

func (uc *UseCase)Logout(user UserModel.User, sid string) error {
	uid, ok := uc.Db.IsOkSession(sid)
	if ok!=nil {
		return ok
	}
	e:=uc.Db.RemoveSession(uid, sid)
	if e!=nil{
		return e
	}
	return nil
}
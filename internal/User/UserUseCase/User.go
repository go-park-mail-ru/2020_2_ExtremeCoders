package UserUseCase

import (
	"CleanArch/internal/User/UserModel"
	"CleanArch/internal/User/UserRepository"
	"CleanArch/internal/errors"
	"net/http"
	"time"
)

type UseCase struct{
	Db UserRepository.UserDB
}

func (uc *UseCase)Signup(user UserModel.User) (error, string) {
	err := uc.Db.IsEmailExists(user.Email)
	if err!=nil{
		return err, ""
	}
	user.Id, err = uc.Db.GenerateUID()
	if err!=nil{
		return err, ""
	}
	sid := string(uc.Db.GenerateSID())
	err = uc.Db.AddUser(&user)
	if err!=nil{
		return err, ""
	}
	err = uc.Db.AddSession(sid, user.Id, &user)
	if err != nil {
		return UserRepository.DbError, ""
	}

	return nil, sid
}

func (uc *UseCase)SignIn(user UserModel.User) (error, *http.Cookie) {
	userEx, erro := uc.Db.GetUserByEmail(user.Email)
	if !erro {
		//return 404, nil
	}
	if userEx.Password != user.Password {
		//return 401, nil
	}
	sid := string(uc.Db.GenerateSID())
	uc.Db.RemoveSessionByUID(userEx.Id)
	if uc.Db.AddSession(sid, userEx.Id, &user) != nil {
		//return 401, nil
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

func (uc *UseCase)Logout(user UserModel.User) (error, *http.Cookie) {
	uid, ok := uc.Db.IsOkSession(session.Value)
	if !ok {
		w.Write(errors.GetErrorWrongCookieAns())
		//glog.Info("RESPONSE: ",getErrorWrongCookieAns())
		return
	}
	e:=uc.Db.RemoveSession(uid, session.Value)
}
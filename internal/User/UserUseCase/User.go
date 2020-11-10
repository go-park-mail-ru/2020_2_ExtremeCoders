package UserUseCase

import (
	"CleanArch/internal/User/UserModel"
	"CleanArch/internal/User/UserRepository"
	"CleanArch/internal/errors"
	err "errors"
	"net/http"
	"time"
)

type UseCase struct {
	Db UserRepository.UserDB
}

var WrongPasswordError = err.New("Wrong password!")

func (uc *UseCase) Signup(user UserModel.User) (error, string) {

	err := uc.Db.IsEmailExists(user.Email)
	if err != nil {
		return err, ""
	}
	user.Id, err = uc.Db.GenerateUID()
	if err != nil {
		return err, ""
	}
	str, err := uc.Db.GenerateSID()
	if err != nil {
		return err, ""
	}
	sid := string(str)
	err = uc.Db.AddUser(&user)
	if err != nil {
		return err, ""
	}
	err = uc.Db.AddSession(sid, user.Id, &user)
	if err != nil {
		return UserRepository.DbError, ""
	}

	return nil, sid
}

func (uc *UseCase) SignIn(user UserModel.User) error {
	userEx, erro := uc.Db.GetUserByEmail(user.Email)
	if erro != nil {
		return erro
	}
	if userEx.Password != user.Password {
		return WrongPasswordError
	}
	sid := string(uc.Db.GenerateSID())
	userToRm, err := uc.Db.GetUserByUID(userEx.Id)
	if err != nil {
		return err
	}
	uc.Db.RemoveSession(userEx.Id)
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

func (uc *UseCase) Logout(user UserModel.User) (error, *http.Cookie) {
	uid, ok := uc.Db.IsOkSession(session.Value)
	if !ok {
		w.Write(errors.GetErrorWrongCookieAns())
		//glog.Info("RESPONSE: ",getErrorWrongCookieAns())
		return
	}
	e := uc.Db.RemoveSession(uid, session.Value)
}

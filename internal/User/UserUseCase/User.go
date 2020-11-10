package UserUseCase

import (
	"CleanArch/internal/User/UserModel"
	"CleanArch/internal/User/UserRepository"
	err "errors"
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

func (uc *UseCase)Logout(sid string) error {
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

func (uc *UseCase)Profile(user UserModel.User) error{
	e :=uc.Db.UpdateProfile(user, user.Email)
	if e !=nil{
		return e
	}
	return nil
}
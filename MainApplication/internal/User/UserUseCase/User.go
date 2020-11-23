package UserUseCase

import (
	"Mailer/MainApplication/internal/User/UserModel"
	"Mailer/MainApplication/internal/User/UserRepository"
	err "errors"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Signup(user UserModel.User) (error, string)
	SignIn(user UserModel.User) (error, string)
	Logout(sid string) error
	Profile(user UserModel.User) error
	GetDB() UserRepository.UserDB
}

type useCase struct {
	Db UserRepository.UserDB
}

func New(db UserRepository.UserDB) UserUseCase {
	return useCase{Db: db}
}

var WrongPasswordError = err.New("Wrong password!")

func (uc useCase) Signup(user UserModel.User) (error, string) {

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

func (uc useCase) SignIn(user UserModel.User) (error, string) {
	userEx, erro := uc.Db.GetUserByEmail(user.Email)
	if erro != nil {
		return erro, ""
	}
	if bcrypt.CompareHashAndPassword([]byte(userEx.Password), []byte(user.Password)) != nil {
		return WrongPasswordError, ""
	}
	sid, e := uc.Db.GenerateSID()
	if e != nil {
		return e, ""
	}
	oldSid, er := uc.Db.GetSessionByUID(userEx.Id)
	if er != nil {
		return er, ""
	}
	er, _ = uc.Db.RemoveSession(oldSid)
	if er != nil {
		return er, ""
	}
	er = uc.Db.AddSession(string(sid), userEx.Id, &user)
	if er != nil {
		return er, ""
	}
	return nil, string(sid)

}

func (uc useCase) Logout(sid string) error {
	_, ok := uc.Db.IsOkSession(sid)
	if ok != nil {
		return ok
	}
	e, _ := uc.Db.RemoveSession(sid)
	if e != nil {
		return e
	}
	return nil
}

func (uc useCase) Profile(user UserModel.User) error {
	e := uc.Db.UpdateProfile(user, user.Email)
	if e != nil {
		return e
	}
	return nil
}

func (uc useCase) GetDB() UserRepository.UserDB {
	return uc.Db
}

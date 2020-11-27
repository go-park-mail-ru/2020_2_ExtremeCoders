package UserUseCase

import (
	"UserService/internal/UserModel"
	proto "UserService/proto"
	"fmt"
)
import "UserService/internal/UserRepository"

type Interface interface {
	IsEmailExists(email *proto.Email) (*proto.Nothing, error)
	AddSession(*proto.AddSessionMsg) (*proto.Nothing, error)
	AddUser(*proto.User) (*proto.Nothing, error)
	GenerateSID(*proto.Nothing) (*proto.Sid, error)
	GenerateUID(*proto.Nothing) (*proto.Uid, error)
	GetUserByEmail(*proto.Email) (*proto.User, error)
	GetUserByUID(*proto.Uid) (*proto.User, error)
	IsOkSession(*proto.Sid) (*proto.Uid, error)
	UpdateProfile(*proto.UpdateProfileMsg) (*proto.Nothing, error)
	RemoveSession(*proto.Sid) (*proto.Uid, error)
	GetSessionByUID(*proto.Uid) (*proto.Sid, error)
}

type UseCase struct {
	db UserRepository.UserDB
}

func New(db UserRepository.UserDB) Interface {
	return UseCase{db: db}
}

func (u UseCase) IsEmailExists(email *proto.Email) (*proto.Nothing, error) {
	err := u.db.IsEmailExists(email.Email)
	return &proto.Nothing{Dummy: true}, err
}

func (u UseCase) AddSession(msg *proto.AddSessionMsg) (*proto.Nothing, error) {
	fmt.Println("ADD SESSION UC", msg.Sid, msg.User.Uid, msg.User.Name)
	userModel := UserModel.User{
		Id:       msg.User.Uid,
		Name:     msg.User.Name,
		Surname:  msg.User.Surname,
		Email:    msg.User.Email,
		Password: msg.User.Password,
	}
	err := u.db.AddSession(msg.Sid, userModel.Id, &userModel)
	return &proto.Nothing{Dummy: true}, err
}

func (u UseCase) AddUser(user *proto.User) (*proto.Nothing, error) {
	userModel := UserModel.User{
		Id:       user.Uid,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: user.Password,
	}
	err := u.db.AddUser(&userModel)
	return &proto.Nothing{Dummy: true}, err
}

func (u UseCase) GenerateSID(nothing *proto.Nothing) (*proto.Sid, error) {
	sid, err := u.db.GenerateSID()
	if err != nil {
		return nil, err
	}
	return &proto.Sid{Sid: string(sid)}, err
}

func (u UseCase) GenerateUID(nothing *proto.Nothing) (*proto.Uid, error) {
	uid, err := u.db.GenerateUID()
	if err != nil {
		return nil, err
	}
	return &proto.Uid{Uid: uid}, err
}

func (u UseCase) GetUserByEmail(email *proto.Email) (*proto.User, error) {
	user, err := u.db.GetUserByEmail(email.Email)
	if err!=nil{
		return nil, err
	}
	res := proto.User{
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Uid:      user.Id,
	}
	return &res, err
}

func (u UseCase) GetUserByUID(uid *proto.Uid) (*proto.User, error) {
	user, err := u.db.GetUserByUID(uid.Uid)
	if err!=nil{
		return nil, err
	}
	res := proto.User{
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Uid:      user.Id,
	}
	return &res, err
}

func (u UseCase) IsOkSession(sid *proto.Sid) (*proto.Uid, error) {
	uid, err := u.db.IsOkSession(sid.Sid)
	if err != nil {
		return nil, err
	}
	return &proto.Uid{Uid: uid}, err
}

func (u UseCase) UpdateProfile(msg *proto.UpdateProfileMsg) (*proto.Nothing, error) {
	userModel := UserModel.User{
		Id:       msg.NewUser.Uid,
		Name:     msg.NewUser.Name,
		Surname:  msg.NewUser.Surname,
		Email:    msg.NewUser.Email,
		Password: msg.NewUser.Password,
	}
	err := u.db.UpdateProfile(userModel, msg.Email)
	if err != nil {
		return nil, err
	}
	return &proto.Nothing{Dummy: true}, err
}

func (u UseCase) RemoveSession(sid *proto.Sid) (*proto.Uid, error) {
	err, uid := u.db.RemoveSession(sid.Sid)
	if err != nil {
		return nil, err
	}
	return &proto.Uid{Uid: uid}, err
}

func (u UseCase) GetSessionByUID(uid *proto.Uid) (*proto.Sid, error) {
	sid, err := u.db.GetSessionByUID(uid.Uid)
	if err != nil {
		return nil, err
	}
	return &proto.Sid{Sid: sid}, err
}

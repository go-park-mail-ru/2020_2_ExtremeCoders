package UserMicroservice

import (
	"Mailer/MainApplication/internal/User/UserModel"
	"Mailer/MainApplication/internal/User/UserRepository"
	userService "Mailer/UserService/proto"
	"context"
)

type UserServiceManager struct {
	usClient userService.UserServiceClient
}

func New(client userService.UserServiceClient) UserRepository.UserDB {
	return UserServiceManager{usClient: client}
}

func (usManager UserServiceManager) IsEmailExists(email string) error {
	ctx := context.Background()
	_, err := usManager.usClient.IsEmailExists(ctx, &userService.Email{Email: email})
	return err
}

func (usManager UserServiceManager) AddUser(user *UserModel.User) error {
	ctx := context.Background()
	_, err := usManager.usClient.AddUser(ctx, &userService.User{
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Uid:      user.Id,
	})
	if err != nil {
		return UserRepository.CantAddUser
	}
	return nil
}

func (usManager UserServiceManager) AddSession(sid string, uid uint64, user *UserModel.User) error {
	ctx := context.Background()
	u := userService.User{
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Uid:      user.Id,
	}
	msg := userService.AddSessionMsg{
		Sid:  sid,
		User: &u,
	}
	_, err := usManager.usClient.AddSession(ctx, &msg)
	if err != nil {
		return UserRepository.CantAddSession
	}
	return nil
}

func (usManager UserServiceManager) GenerateSID() ([]rune, error) {
	ctx := context.Background()
	sid, err := usManager.usClient.GenerateSID(ctx, &userService.Nothing{Dummy: true})
	if err != nil {
		return nil, err
	}
	return []rune(sid.Sid), nil
}

func (usManager UserServiceManager) GenerateUID() (uint64, error) {
	ctx := context.Background()
	uid, err := usManager.usClient.GenerateUID(ctx, &userService.Nothing{Dummy: true})
	if err != nil {
		return 0, err
	}
	return uid.Uid, nil
}

func (usManager UserServiceManager) GetUserByEmail(email string) (*UserModel.User, error) {
	ctx := context.Background()
	user, err := usManager.usClient.GetUserByEmail(ctx, &userService.Email{Email: email})
	if err != nil {
		return nil, UserRepository.CantGetUserByEmail
	}
	u := UserModel.User{
		Id:       user.Uid,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: user.Password,
	}
	return &u, nil
}

func (usManager UserServiceManager) GetUserByUID(uid uint64) (*UserModel.User, error) {
	ctx := context.Background()
	user, err := usManager.usClient.GetUserByUID(ctx, &userService.Uid{Uid: uid})
	if err != nil {
		return nil, UserRepository.CantGetUserByUid
	}
	u := UserModel.User{
		Id:       user.Uid,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: user.Password,
	}
	return &u, nil
}

func (usManager UserServiceManager) IsOkSession(sid string) (uint64, error) {
	ctx := context.Background()
	uid, err := usManager.usClient.IsOkSession(ctx, &userService.Sid{Sid: sid})
	if err != nil {
		return 0, UserRepository.InvalidSession
	}
	return uid.Uid, nil
}

func (usManager UserServiceManager) UpdateProfile(newUser UserModel.User, email string) error {
	ctx := context.Background()
	u := userService.User{
		Email:    newUser.Email,
		Name:     newUser.Name,
		Surname:  newUser.Surname,
		Password: newUser.Password,
		Uid:      newUser.Id,
	}
	msg := userService.UpdateProfileMsg{
		NewUser: &u,
		Email:   email,
	}
	_, err := usManager.usClient.UpdateProfile(ctx, &msg)
	if err != nil {
		return UserRepository.CantUpdateUser
	}
	return nil
}

func (usManager UserServiceManager) RemoveSession(sid string) (error, uint64) {
	ctx := context.Background()
	uid, err := usManager.usClient.RemoveSession(ctx, &userService.Sid{Sid: sid})
	if err != nil {
		return UserRepository.RemoveSessionError, 0
	}
	return nil, uid.Uid
}

func (usManager UserServiceManager) GetSessionByUID(uid uint64) (string, error) {
	ctx := context.Background()
	sid, err := usManager.usClient.GetSessionByUID(ctx, &userService.Uid{Uid: uid})
	if err != nil {
		return "", UserRepository.GetSessionError
	}
	return sid.Sid, nil
}

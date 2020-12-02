package test

import (
	"MainApplication/internal/User/UserModel"
	"MainApplication/internal/User/UserRepository/UserMicroservice"
	userService "MainApplication/proto/UserServise"
	mock "MainApplication/test/mock_UserProto"
	"context"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestIsEmailExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()
	mockLetter.EXPECT().IsEmailExists(ctx, &userService.Email{ Email: user.Email}).Times(1)
	uc := UserMicroservice.New(mockLetter)

	uc.IsEmailExists(user.Email)
}

func TestAddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()
	mockLetter.EXPECT().AddUser(ctx, &userService.User{ Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Uid:      user.Id,}).Times(1)
	uc := UserMicroservice.New(mockLetter)

	uc.AddUser(&user)
}

func TestAddSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()
	u := userService.User{
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Uid:      user.Id,
	}
	sid:="asjhdflashdbfp"
	msg := userService.AddSessionMsg{
		Sid:  sid,
		User: &u,
	}
	mockLetter.EXPECT().AddSession(ctx, &msg).Times(1)
	uc := UserMicroservice.New(mockLetter)

	uc.AddSession(sid, user.Id,&user)
}

func TestGetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()
	mockLetter.EXPECT().GetUserByEmail(ctx, &userService.Email{Email: user.Email}).Return(user, nil)
	uc := UserMicroservice.New(mockLetter)

	uc.GetUserByEmail(user.Email)
}
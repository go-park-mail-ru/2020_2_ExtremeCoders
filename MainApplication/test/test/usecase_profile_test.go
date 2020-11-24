package test

import (
	"CleanArch/MainApplication/internal/User/UserModel"
	"CleanArch/MainApplication/internal/User/UserRepository"
	"CleanArch/MainApplication/internal/User/UserUseCase"
	mock "CleanArch/MainApplication/test/mock_UserRepository"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().UpdateProfile(user, user.Email).Return(nil)
	uc := UserUseCase.New(mockLetter)
	uc.Profile(user)
}

func TestProfileUpUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().UpdateProfile(user, user.Email).Return(UserRepository.CantUpdateUser)
	uc := UserUseCase.New(mockLetter)
	uc.Profile(user)
}
package test

import (
	mock "2020_2_ExtremeCoders/MainApplication/test/mock_UserRepository"
	"2020_2_ExtremeCoders/internal/User/UserModel"
	"CleanArch/internal/User/UserRepository"
	"CleanArch/internal/User/UserUseCase"
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
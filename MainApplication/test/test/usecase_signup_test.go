package test

import (
	"MainApplication/internal/User/UserModel"
	"MainApplication/internal/User/UserRepository"
	"MainApplication/internal/User/UserUseCase"
	mock "MainApplication/test/mock_UserRepository"

	"github.com/golang/mock/gomock"
	"testing"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	var sid []rune
	sid = []rune("VLbutPK_aMA_zVi4QP_EL_7KLXl8Uxwg")
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().IsEmailExists(user.Email).Return(nil)
	mockLetter.EXPECT().GenerateUID().Return(user.Id, nil)
	mockLetter.EXPECT().GenerateSID().Return(sid, nil)
	mockLetter.EXPECT().AddUser(&user).Return(nil)
	mockLetter.EXPECT().AddSession(string(sid), uint64(user.Id), &user).Return(nil)
	uc := UserUseCase.New(mockLetter)

	uc.Signup(user)
}

func TestSaveLetterExEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().IsEmailExists(user.Email).Return(UserRepository.EmailAlreadyExists)

	uc := UserUseCase.New(mockLetter)

	uc.Signup(user)
}

func TestSaveLetterGenUID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().IsEmailExists(user.Email).Return(nil)
	mockLetter.EXPECT().GenerateUID().Return(user.Id, UserRepository.InvalidSession)
	uc := UserUseCase.New(mockLetter)
	uc.Signup(user)
}

func TestSaveLetterGenSID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().IsEmailExists(user.Email).Return(nil)
	mockLetter.EXPECT().GenerateUID().Return(user.Id, nil)
	mockLetter.EXPECT().GenerateSID().Return([]rune(""), UserRepository.InvalidSession)
	uc := UserUseCase.New(mockLetter)
	uc.Signup(user)
}

func TestSaveLetterAddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	var sid []rune
	sid = []rune("VLbutPK_aMA_zVi4QP_EL_7KLXl8Uxwg")
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().IsEmailExists(user.Email).Return(nil)
	mockLetter.EXPECT().GenerateUID().Return(user.Id, nil)
	mockLetter.EXPECT().GenerateSID().Return(sid, nil)
	mockLetter.EXPECT().AddUser(&user).Return(UserRepository.CantAddUser)

	uc := UserUseCase.New(mockLetter)

	uc.Signup(user)
}

func TestSaveLetterAddSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	var sid []rune
	sid = []rune("VLbutPK_aMA_zVi4QP_EL_7KLXl8Uxwg")
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().IsEmailExists(user.Email).Return(nil)
	mockLetter.EXPECT().GenerateUID().Return(user.Id, nil)
	mockLetter.EXPECT().GenerateSID().Return(sid, nil)
	mockLetter.EXPECT().AddUser(&user).Return(nil)
	mockLetter.EXPECT().AddSession(string(sid), uint64(user.Id), &user).Return(UserRepository.CantAddSession)
	uc := UserUseCase.New(mockLetter)

	uc.Signup(user)
}

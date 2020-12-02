package test

import (
	"MainApplication/internal/User/UserModel"
	"MainApplication/internal/User/UserRepository/UserMicroservice"
	mock "MainApplication/test/mock_UserRepository"
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
	var sid []rune
	sid=[]rune("VLbutPK_aMA_zVi4QP_EL_7KLXl8Uxwg")
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().IsEmailExists(user.Email).Return(nil)
	mockLetter.EXPECT().GenerateUID().Return(user.Id,nil)
	mockLetter.EXPECT().GenerateSID().Return(sid, nil)
	mockLetter.EXPECT().AddUser(&user).Return(nil)
	mockLetter.EXPECT().AddSession(string(sid), uint64(user.Id),&user).Return(nil)
	uc := UserUseCase.New(mockLetter)

	uc.Signup(user)
}
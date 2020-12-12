package test

import (
	"Mailer/MainApplication/internal/User/UserModel"
	"Mailer/MainApplication/internal/User/UserRepository"
	"Mailer/MainApplication/internal/User/UserUseCase"
	mock "Mailer/MainApplication/test/mock_UserRepository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	userex := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "$2a$14$OzJS/7LjHhx8U8vh6/hl5uPx3X2OGhrRHNYalvAHXaF9Ko8Uooef.",
	}
	var sid []rune
	sid=[]rune("VLbutPK_aMA_zVi4QP_EL_7KLXl8Uxwg")
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().GetUserByEmail(user.Email).Return(&userex,nil)
	mockLetter.EXPECT().GenerateSID().Return(sid,nil)
	mockLetter.EXPECT().GetSessionByUID(user.Id).Return(string(sid),nil)
	mockLetter.EXPECT().RemoveSession(string(sid)).Return(nil,uint64(0))
	mockLetter.EXPECT().AddSession(string(sid), user.Id, &userex).Return(nil)
	uc := UserUseCase.New(mockLetter)

	err, _:=uc.SignIn(user)
	assert.Nil(t, err)

}

func TestSignInByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	userex := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "$2a$14$OzJS/7LjHhx8U8vh6/hl5uPx3X2OGhrRHNYalvAHXaF9Ko8Uooef.",
	}
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().GetUserByEmail(user.Email).Return(&userex,UserRepository.CantGetUserByEmail)
	uc := UserUseCase.New(mockLetter)

	uc.SignIn(user)

}

func TestSignInGenSid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	userex := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "$2a$14$OzJS/7LjHhx8U8vh6/hl5uPx3X2OGhrRHNYalvAHXaF9Ko8Uooef.",
	}
	var sid []rune
	sid=[]rune("VLbutPK_aMA_zVi4QP_EL_7KLXl8Uxwg")
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().GetUserByEmail(user.Email).Return(&userex,nil)
	mockLetter.EXPECT().GenerateSID().Return(sid,UserRepository.InvalidSession)
	uc := UserUseCase.New(mockLetter)

	uc.SignIn(user)
}

func TestSignInSessByUID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	userex := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "$2a$14$OzJS/7LjHhx8U8vh6/hl5uPx3X2OGhrRHNYalvAHXaF9Ko8Uooef.",
	}
	var sid []rune
	sid=[]rune("VLbutPK_aMA_zVi4QP_EL_7KLXl8Uxwg")
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().GetUserByEmail(user.Email).Return(&userex,nil)
	mockLetter.EXPECT().GenerateSID().Return(sid,nil)
	mockLetter.EXPECT().GetSessionByUID(user.Id).Return(string(sid),UserRepository.InvalidSession)
	uc := UserUseCase.New(mockLetter)
	uc.SignIn(user)
}

func TestSignInRemSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	userex := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "$2a$14$OzJS/7LjHhx8U8vh6/hl5uPx3X2OGhrRHNYalvAHXaF9Ko8Uooef.",
	}
	var sid []rune
	sid=[]rune("VLbutPK_aMA_zVi4QP_EL_7KLXl8Uxwg")
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().GetUserByEmail(user.Email).Return(&userex,nil)
	mockLetter.EXPECT().GenerateSID().Return(sid,nil)
	mockLetter.EXPECT().GetSessionByUID(user.Id).Return(string(sid),nil)
	mockLetter.EXPECT().RemoveSession(string(sid)).Return(UserRepository.RemoveSessionError,uint64(0))
	uc := UserUseCase.New(mockLetter)
	uc.SignIn(user)
}

func TestSignInAddSess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	userex := UserModel.User{
		Id: 123,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "$2a$14$OzJS/7LjHhx8U8vh6/hl5uPx3X2OGhrRHNYalvAHXaF9Ko8Uooef.",
	}
	var sid []rune
	sid=[]rune("VLbutPK_aMA_zVi4QP_EL_7KLXl8Uxwg")
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().GetUserByEmail(user.Email).Return(&userex,nil)
	mockLetter.EXPECT().GenerateSID().Return(sid,nil)
	mockLetter.EXPECT().GetSessionByUID(user.Id).Return(string(sid),nil)
	mockLetter.EXPECT().RemoveSession(string(sid)).Return(nil,uint64(0))
	mockLetter.EXPECT().AddSession(string(sid), user.Id, &userex).Return(UserRepository.CantAddSession)
	uc := UserUseCase.New(mockLetter)
	uc.SignIn(user)
}
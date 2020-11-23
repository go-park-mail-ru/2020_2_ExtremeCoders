package test

import (
	mock "2020_2_ExtremeCoders/MainApplication/test/mock_UserRepository"
	"2020_2_ExtremeCoders/internal/User/UserRepository"
	"CleanArch/internal/User/UserUseCase"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var sid []rune
	sid=[]rune("VLbutPK_aMA_zVi4QP_EL_7KLXl8Uxwg")
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().IsOkSession(string(sid)).Return(uint64(0),nil)
	mockLetter.EXPECT().RemoveSession(string(sid)).Return(nil,uint64(0))
	uc := UserUseCase.New(mockLetter)
	uc.Logout(string(sid))
}

func TestLogoutIsOkSess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var sid []rune
	sid=[]rune("VLbutPK_aMA_zVi4QP_EL_7KLXl8Uxwg")
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().IsOkSession(string(sid)).Return(uint64(0), UserRepository.GetSessionError)
	uc := UserUseCase.New(mockLetter)
	uc.Logout(string(sid))
}

func TestLogoutRemoveSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var sid []rune
	sid=[]rune("VLbutPK_aMA_zVi4QP_EL_7KLXl8Uxwg")
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().IsOkSession(string(sid)).Return(uint64(0),nil)
	mockLetter.EXPECT().RemoveSession(string(sid)).Return(UserRepository.RemoveSessionError,uint64(0))
	uc := UserUseCase.New(mockLetter)
	uc.Logout(string(sid))
}
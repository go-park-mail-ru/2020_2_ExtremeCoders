package test

import (
	"CleanArch/internal/User/UserDelivery"
	"CleanArch/internal/errors"
	mock "CleanArch/test/mock_UserUseCase"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)


func TestUserDelivery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUseCase := mock.MockUserUseCase{}
	userDelivery := UserDelivery.New(&mockUseCase)
	mockUseCase.EXPECT().SignIn(nil).MaxTimes(0)
	w := MyWriter{}
	request := http.Request{}
	userDelivery.SignIn(&w, &request)
	if string(w.Str)!=string(errors.GetErrorUnexpectedAns()){
		t.Errorf("Expected error is  " + string(errors.GetErrorUnexpectedAns()))
	}
}

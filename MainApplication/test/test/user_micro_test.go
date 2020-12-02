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
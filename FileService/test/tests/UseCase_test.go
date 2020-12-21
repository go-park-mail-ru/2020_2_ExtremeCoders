package test

import (
"Mailer/FileService/internal/File/UseCase"
mock "Mailer/FileService/test/mock_Repository"
proto "Mailer/FileService/proto"
	"errors"
	"github.com/golang/mock/gomock"
"testing"
)

func TestGetAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	user := &proto.User{Email: ""}
	mockRepo := mock.NewMockInterface(ctrl)
	mockRepo.EXPECT().GetAvatar(user).Return(&proto.Avatar{}, errors.New("sdf"))
	mockRepo.EXPECT().GetDefaultAvatar().Times(1)
	uc := UseCase.New(mockRepo)
	uc.GetAvatar(user)
}

func TestGetFiles(t *testing.T)  {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	lid := &proto.LetterId{Id: 1}
	mockRepo := mock.NewMockInterface(ctrl)
	mockRepo.EXPECT().GetFiles(lid)
	uc := UseCase.New(mockRepo)
	uc.GetFiles(lid)
}

func TestSaveAvatar(t *testing.T)  {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	avatar := &proto.Avatar{Email:"email"}
	mockRepo := mock.NewMockInterface(ctrl)
	mockRepo.EXPECT().SaveAvatar(avatar).Times(1)
	uc := UseCase.New(mockRepo)
	uc.SaveAvatar(avatar)
}

func Test(t *testing.T)  {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	file:= &proto.Files{}
	mockRepo := mock.NewMockInterface(ctrl)
	mockRepo.EXPECT().SaveFiles(file).Times(1)
	uc := UseCase.New(mockRepo)
	uc.SaveFiles(file)
}

package test

import (
"FileService/internal/File/UseCase"
mock "FileService/test/mock_Repository"
proto "FileService/proto"
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


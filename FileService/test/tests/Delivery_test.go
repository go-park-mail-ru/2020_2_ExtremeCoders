package test

import (
	"Mailer/FileService/internal/File/Delivery"
	proto "Mailer/FileService/proto"
	mock "Mailer/FileService/test/mock_UseCase"
	"context"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestSaveFilesDelivery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	files := &proto.Files{}
	mockUseCase := mock.NewMockInterface(ctrl)
	mockUseCase.EXPECT().SaveFiles(files).Times(1)
	de := Delivery.NewFileManager(mockUseCase)
	ctx:=context.Background()
	de.SaveFiles(ctx, files)
}

func TestGetFilesDelivery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	lid := &proto.LetterId{}
	mockUseCase := mock.NewMockInterface(ctrl)
	mockUseCase.EXPECT().GetFiles(lid).Times(1)
	de := Delivery.NewFileManager(mockUseCase)
	ctx:=context.Background()
	de.GetFiles(ctx, lid)
}

func TestGetAvatarDelivery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	user := &proto.User{}
	mockUseCase := mock.NewMockInterface(ctrl)
	mockUseCase.EXPECT().GetAvatar(user).Times(1)
	de := Delivery.NewFileManager(mockUseCase)
	ctx:=context.Background()
	de.GetAvatar(ctx, user)
}

func TestSetAvatarDelivery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	avatar := &proto.Avatar{}
	mockUseCase := mock.NewMockInterface(ctrl)
	mockUseCase.EXPECT().SaveAvatar(avatar).Times(1)
	de := Delivery.NewFileManager(mockUseCase)
	ctx:=context.Background()
	de.SetAvatar(ctx, avatar)
}


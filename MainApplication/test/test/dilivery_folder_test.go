package test

import (
	"Mailer/MainApplication/internal/Folder/FolderDelivery"
	mockMail "Mailer/MainApplication/test/mock_MailServiceProto"
	mockUser "Mailer/MainApplication/test/mock_UserProto"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/url"
	"testing"
)

type MyWriter1 struct {
	Str []byte
}

func (writer MyWriter1) Header() http.Header {
	fmt.Println("implement me")
	return nil
}

func (writer *MyWriter1) Write(bytes []byte) (int, error) {
	writer.Str = append(writer.Str, bytes...)
	return len(writer.Str), nil
}

func (writer MyWriter1) WriteHeader(statusCode int) {
	fmt.Println("implement me")
}

func TestRenameFolder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := http.Request{Method: "POST", URL: &url.URL{
		Path: "/user/folders/sended/folderName",
	}}
	w := MyWriter1{}

	mailClient := mockMail.NewMockLetterServiceClient(ctrl)
	userClient := mockUser.NewMockUserServiceClient(ctrl)
	userClient.EXPECT().RenameFolder(r, w).Times(0)
	fd := FolderDelivery.New(userClient, mailClient)
	fd.RenameFolder(&w, &r)
}

func TestRemoveFolder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := http.Request{Method: "POST", URL: &url.URL{
		Path: "/user/folders/sended/folderName",
	}}
	w := MyWriter1{}
	mailClient := mockMail.NewMockLetterServiceClient(ctrl)
	userClient := mockUser.NewMockUserServiceClient(ctrl)
	userClient.EXPECT().RemoveFolder(r, w).Times(0)
	fd := FolderDelivery.New(userClient, mailClient)
	fd.RemoveFolder(&w, &r)
}

func Test(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := http.Request{Method: "POST", URL: &url.URL{
		Path: "/user/folders/sended/folderName",
	}}
	w := MyWriter1{}
	mailClient := mockMail.NewMockLetterServiceClient(ctrl)
	userClient := mockUser.NewMockUserServiceClient(ctrl)
	userClient.EXPECT().GetFolderId(r, w).Times(0)
	mailClient.EXPECT().RemoveLetterFromDir(r, w).Times(0)
	fd := FolderDelivery.New(userClient, mailClient)
	fd.RemoveLetterInFolder(&w, &r)
}

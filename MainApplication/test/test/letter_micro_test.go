package test

import (
	Model "MainApplication/internal/Letter/LetterModel"
	"MainApplication/internal/Letter/LetterRepository/LetterService"
	letterService "MainApplication/proto/MailService"
	mock "MainApplication/test/mock_MailServiceProto"
	"context"
	"github.com/golang/mock/gomock"
	"testing"
)

var lettermodel = Model.Letter{
	Id: 123,
	Sender: "Dellvin",
	Receiver: "Black",
}

var letter = letterService.Letter{
Lid: 123,
Sender: "Dellvin",
Receiver: "Black",
}

var response=letterService.Response{Description: "ok", Ok: true}

var lr =letterService.LetterResponse{
	Letter: &letter,
	Result: &response,
}

var letterListResponse letterService.LetterListResponse


func TestWatchLetter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	letter := letterService.Letter{
		Lid: 123,
		Sender: "Dellvin",
		Receiver: "Black",
	}
	var lr letterService.LetterResponse
	lr.Letter=&letter
	mockLetter := mock.NewMockLetterServiceClient(ctrl)
	ctx := context.Background()

	var lid uint64 =345

	mockLetter.EXPECT().WatchedLetter(ctx, &letterService.Lid{Lid: uint64(lid)}).Return(&lr, nil)
	uc := LetterService.New(mockLetter)

	uc.WatchLetter(lid)
}


func TestSaveMail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()


	mockLetter := mock.NewMockLetterServiceClient(ctrl)
	ctx := context.Background()



	mockLetter.EXPECT().SaveLetter(ctx, &letter).Return(&response, nil)
	uc := LetterService.New(mockLetter)

	uc.SaveMail(lettermodel)
}

func TestGetReceivedLetters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()


	mockLetter := mock.NewMockLetterServiceClient(ctrl)
	ctx := context.Background()

	email:="kadsf"
	letterListResponse.Letter=append(letterListResponse.Letter, &letter)
	letterListResponse.Result=&response
	mockLetter.EXPECT().GetLettersRecv(ctx, &letterService.Email{Email: email}).Return(&letterListResponse, nil)
	uc := LetterService.New(mockLetter)

	uc.GetReceivedLetters(email)
}


func TestGetSendedLetters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()


	mockLetter := mock.NewMockLetterServiceClient(ctrl)
	ctx := context.Background()

	email:="kadsf"
	letterListResponse.Letter=append(letterListResponse.Letter, &letter)
	letterListResponse.Result=&response
	mockLetter.EXPECT().GetLettersSend(ctx, &letterService.Email{Email: email}).Return(&letterListResponse, nil)
	uc := LetterService.New(mockLetter)

	uc.GetSendedLetters(email)
}

func TestGetReceivedLettersDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()


	mockLetter := mock.NewMockLetterServiceClient(ctrl)
	ctx := context.Background()


	dir:=1232342
	letterListResponse.Letter=append(letterListResponse.Letter, &letter)
	letterListResponse.Result=&response
	mockLetter.EXPECT().GetLettersByDirRecv(ctx, &letterService.DirName{DirName: uint64(dir)}).Return(&letterListResponse, nil)
	uc := LetterService.New(mockLetter)

	uc.GetReceivedLettersDir(uint64(dir))
}

func TestGetSendedLettersDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()


	mockLetter := mock.NewMockLetterServiceClient(ctrl)
	ctx := context.Background()


	dir:=1232342
	letterListResponse.Letter=append(letterListResponse.Letter, &letter)
	letterListResponse.Result=&response
	mockLetter.EXPECT().GetLettersByDirSend(ctx, &letterService.DirName{DirName: uint64(dir)}).Return(&letterListResponse, nil)
	uc := LetterService.New(mockLetter)

	uc.GetSendedLettersDir(uint64(dir))
}
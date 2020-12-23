package test

import (
	letterService "Mailer/MailService/proto"
	Model "Mailer/MainApplication/internal/Letter/LetterModel"
	"Mailer/MainApplication/internal/Letter/LetterRepository/LetterService"
	mock "Mailer/MailService/test/mock_MailServiceProto"
	"context"
	"github.com/golang/mock/gomock"
	"testing"
)

var lettermodel = Model.Letter{
	Id:       123,
	Sender:   "Dellvin",
	Receiver: "Black",
}

var letter = letterService.Letter{
	Lid:      123,
	Sender:   "Dellvin",
	Receiver: "Black",
}

var response = letterService.Response{Description: "ok", Ok: true}

var letterListResponse letterService.LetterListResponse

func TestWatchLetter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var lr letterService.LetterResponse
	lr.Letter = &letter
	mockLetter := mock.NewMockLetterServiceClient(ctrl)
	ctx := context.Background()

	var lid uint64 = 345

	mockLetter.EXPECT().WatchedLetter(ctx, &letterService.Lid{Lid: uint64(lid)}).Return(&lr, nil)
	uc := LetterService.New(mockLetter)

	_, _ = uc.WatchLetter(lid)
}

func TestSaveMail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterServiceClient(ctrl)
	ctx := context.Background()

	mockLetter.EXPECT().SaveLetter(ctx, &letter).Return(&response, nil)
	uc := LetterService.New(mockLetter)

	_ = uc.SaveMail(lettermodel)
}

func TestGetReceivedLetters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterServiceClient(ctrl)
	ctx := context.Background()

	email := "kadsf"
	letterListResponse.Letter = append(letterListResponse.Letter, &letter)
	letterListResponse.Result = &response
	mockLetter.EXPECT().GetLettersRecv(ctx, &letterService.Email{Email: email}).Return(&letterListResponse, nil)
	uc := LetterService.New(mockLetter)

	_, _ = uc.GetReceivedLetters(email, 0, 0)
}

func TestGetSendedLettersewr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterServiceClient(ctrl)
	ctx := context.Background()

	email := "kadsf"
	letterListResponse.Letter = append(letterListResponse.Letter, &letter)
	letterListResponse.Result = &response
	mockLetter.EXPECT().GetLettersSend(ctx, &letterService.Email{Email: email,
		Limit: uint64(5), Offset: uint64(0)}).Return(&letterListResponse, nil)
	uc := LetterService.New(mockLetter)

	_, _ = uc.GetSendedLetters(email, uint64(5), uint64(0))
}

func TestGetReceivedLettersDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterServiceClient(ctrl)
	ctx := context.Background()

	dir := 1232342
	letterListResponse.Letter = append(letterListResponse.Letter, &letter)
	letterListResponse.Result = &response
	mockLetter.EXPECT().GetLettersByDirRecv(ctx, &letterService.DirName{DirName: uint64(dir),
		Limit: uint64(5), Offset: uint64(0)}).Return(&letterListResponse, nil)
	uc := LetterService.New(mockLetter)

	_, _ = uc.GetReceivedLettersDir(uint64(dir), uint64(5), uint64(0))
}

func TestGetSendedLettersDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterServiceClient(ctrl)
	ctx := context.Background()

	dir := 1232342
	letterListResponse.Letter = append(letterListResponse.Letter, &letter)
	letterListResponse.Result = &response
	mockLetter.EXPECT().GetLettersByDirSend(ctx, &letterService.DirName{DirName: uint64(dir)}).Return(&letterListResponse, nil)
	uc := LetterService.New(mockLetter)

	_, _ = uc.GetSendedLettersDir(uint64(dir))
}

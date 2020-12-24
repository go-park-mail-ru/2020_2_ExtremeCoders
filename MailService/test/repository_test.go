package test

import (
	"Mailer/MailService/internal/UseCase"
	"Mailer/MailService/internal/Model"
	mock "Mailer/MailService/test/mock_LetterRepository"
	"github.com/golang/mock/gomock"
	"testing"
)

var Letter = Model.Letter{Receiver: "dellvin.black@gmail.com"}

func TestGetLettersRecvDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetLettersRecvDir(uint64(0),uint64(0),uint64(0)).Return(nil, nil)
	uc := UseCase.New(mockLetter)

	uc.GetLettersRecvDir(uint64(0),uint64(0),uint64(0))
}

func TestGetLettersSendDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetLettersSentDir(uint64(0)).Return(nil, nil)
	uc := UseCase.New(mockLetter)

	uc.GetLettersSendDir(0)
}

//func TestSaveLetter(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockLetter := mock.NewMockLetterDB(ctrl)
//	mockLetter.EXPECT().SaveMail(Letter).Return(nil)
//	uc := UseCase.New(mockLetter)
//
//	uc.SaveLetter(Letter)
//}

func TestWatchLetter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().SetLetterWatched(uint64(0)).Return(nil, Letter)
	uc := UseCase.New(mockLetter)

	uc.WatchLetter(0)
}


func TestGetLettersSend(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetLettersSent(Letter.Receiver,uint64(0),uint64(0)).Return(nil, nil)
	uc := UseCase.New(mockLetter)

	uc.GetLettersSend(Letter.Receiver,uint64(0),uint64(0))
}


func TestGetLettersRecv(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetLettersRecv(Letter.Receiver,uint64(0),uint64(0)).Return(nil, nil)
	uc := UseCase.New(mockLetter)

	uc.GetLettersRecv(Letter.Receiver,uint64(0),uint64(0))
}

func TestAddLetterToDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().AddLetterToDir(uint64(0), uint64(0), true).Return(nil)
	uc := UseCase.New(mockLetter)

	uc.AddLetterToDir(uint64(0), uint64(0), true)
}

func TestRemoveLetterToDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().RemoveLetterFromDir(uint64(0), uint64(0), true).Return(nil)
	uc := UseCase.New(mockLetter)

	uc.RemoveLetterFromDir(uint64(0), uint64(0), true)
}

func TestRemoveDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().RemoveDir(uint64(0), true).Return(nil)
	uc := UseCase.New(mockLetter)

	uc.RemoveDir(uint64(0), true)
}
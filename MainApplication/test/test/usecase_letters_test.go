package test

import (
	"Mailer/MainApplication/internal/Letter/LetterModel"
	"Mailer/MainApplication/internal/Letter/LetterUseCase"
	mock "Mailer/MainApplication/test/mock_LetterRepository"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetSendedLetters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetSendedLetters("dellvin.black@gmail.com", uint64(5), uint64(0)).Return(nil, nil)
	uc := LetterUseCase.New(mockLetter)
	_, _ = uc.GetSendedLetters("dellvin.black@gmail.com", uint64(5), uint64(0))
}

func TestGetRecivedLetters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetReceivedLetters("dellvin.black@gmail.com", uint64(1), uint64(0)).Return(nil, nil)
	uc := LetterUseCase.New(mockLetter)
	_, _ = uc.GetReceivedLetters("dellvin.black@gmail.com", uint64(1), uint64(0))
}

func TestGetRecivedLettersDir(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetReceivedLettersDir(uint64(12), uint64(5), uint64(0)).Return(nil, nil)
	uc := LetterUseCase.New(mockLetter)
	_, _ = uc.GetReceivedLettersDir(uint64(12), uint64(5), uint64(0))
}

func TestGetSendedLettersDirasdf(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetSendedLettersDir(uint64(12)).Return(nil, nil)
	uc := LetterUseCase.New(mockLetter)
	_, _ = uc.GetSendedLettersDir(uint64(12))
}

func TestSaveMailasdf(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLetter := mock.NewMockLetterDB(ctrl)
	l := LetterModel.Letter{}
	mockLetter.EXPECT().SaveMail(l).Return(nil)
	uc := LetterUseCase.New(mockLetter)
	_ = uc.SaveLetter(&l)
}

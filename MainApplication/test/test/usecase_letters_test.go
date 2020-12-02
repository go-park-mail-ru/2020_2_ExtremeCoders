package test

import (
	"MainApplication/internal/Letter/LetterUseCase"
	mock "MainApplication/test/mock_LetterRepository"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetSendedLetters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetSendedLetters("dellvin.black@gmail.com").Return(nil, nil)
	uc := LetterUseCase.New(mockLetter)
	uc.GetSendedLetters("dellvin.black@gmail.com")
}

func TestGetRecivedLetters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetReceivedLetters("dellvin.black@gmail.com").Return(nil, nil)
	uc := LetterUseCase.New(mockLetter)
	uc.GetReceivedLetters("dellvin.black@gmail.com")
}

//func TestGetRecivedLettersDir(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//	mockLetter := mock.NewMockLetterDB(ctrl)
//	mockLetter.EXPECT().GetReceivedLettersDir(12).Return(nil, nil)
//	uc := LetterUseCase.New(mockLetter)
//	uc.GetReceivedLettersDir(12)
//}
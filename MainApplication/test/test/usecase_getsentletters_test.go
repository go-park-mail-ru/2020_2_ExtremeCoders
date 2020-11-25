package test

import (
	mock "2020_2_ExtremeCoders/MainApplication/test/mock_LetterRepository"
	"2020_2_ExtremeCoders/internal/Letter/LetterRepository"
	"CleanArch/internal/Letter/LetterUseCase"
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

func TestSent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetSendedLetters("dellvin.black@gmail.com").Return(LetterRepository.ReceiverNotFound, nil)
	uc := LetterUseCase.New(mockLetter)
	uc.GetSendedLetters("dellvin.black@gmail.com")
}

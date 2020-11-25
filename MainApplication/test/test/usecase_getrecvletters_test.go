package test

import (
	mock "2020_2_ExtremeCoders/MainApplication/test/mock_LetterRepository"
	"2020_2_ExtremeCoders/internal/Letter/LetterRepository"
	"CleanArch/internal/Letter/LetterUseCase"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetReceivedLetters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetReceivedLetters("dellvin.black@gmail.com").Return(nil, nil)

	uc := LetterUseCase.New(mockLetter)

	uc.GetReceivedLetters("dellvin.black@gmail.com")
}

func TestRecv(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetReceivedLetters("dellvin.black@gmail.com").Return(LetterRepository.ReceivedLetterError, nil)

	uc := LetterUseCase.New(mockLetter)

	uc.GetReceivedLetters("dellvin.black@gmail.com")
}

package test

import (
	mock "MailService/test/mock_LetterRepository"
	"MainApplication/internal/Letter/LetterRepository"
	"MainApplication/internal/Letter/LetterUseCase"
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
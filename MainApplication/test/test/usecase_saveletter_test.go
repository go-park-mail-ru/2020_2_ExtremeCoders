package test

import (
	"MainApplication/internal/Letter/LetterModel"
	"MainApplication/internal/Letter/LetterRepository"
	"MainApplication/internal/Letter/LetterUseCase"
	mock "MainApplication/test/mock_LetterRepository"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestSaveLetter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	Letter := &LetterModel.Letter{Receiver: "dellvin.black@gmail.com"}
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GenerateLID().Return(uint64(0))
	mockLetter.EXPECT().IsUserExist((*Letter).Receiver).Return(nil)
	mockLetter.EXPECT().SaveMail(*Letter).Return(nil)
	uc := LetterUseCase.New(mockLetter)

	uc.SaveLetter(Letter)
}

func TestIsUserExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	Letter := &LetterModel.Letter{Receiver: "dellvin.black@gmail.com"}
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GenerateLID().Return(uint64(0))
	mockLetter.EXPECT().IsUserExist((*Letter).Receiver).Return(LetterRepository.ReceiverNotFound)

	uc := LetterUseCase.New(mockLetter)

	uc.SaveLetter(Letter)
}

func TestSaveMail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	Letter := &LetterModel.Letter{Receiver: "dellvin.black@gmail.com"}
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GenerateLID().Return(uint64(0))
	mockLetter.EXPECT().IsUserExist((*Letter).Receiver).Return(nil)
	mockLetter.EXPECT().SaveMail(*Letter).Return(LetterRepository.SaveLetterError)
	uc := LetterUseCase.New(mockLetter)

	uc.SaveLetter(Letter)
}

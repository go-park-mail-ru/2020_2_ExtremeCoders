package test

import (
	"CleanArch/internal/Letter/LetterModel"
	"CleanArch/internal/Letter/LetterUseCase"
	mock "CleanArch/test/letter/mock_LetterRepository"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestSaveLetter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	Letter:=&LetterModel.Letter{Receiver: "dellvin.black@gmail.com"}
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GenerateLID().Return(uint64(0))
	mockLetter.EXPECT().IsUserExist((*Letter).Receiver).Return(nil)
	mockLetter.EXPECT().SaveMail(*Letter).Return(nil)
	uc:=LetterUseCase.UseCase{Db: mockLetter}

	uc.SaveLetter(Letter)
}

func TestGetReceivedLetters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetReceivedLetters("dellvin.black@gmail.com").Return(nil, nil)

	uc:=LetterUseCase.UseCase{Db: mockLetter}

	uc.GetReceivedLetters("dellvin.black@gmail.com")
}

func TestGetSendedLetters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLetter := mock.NewMockLetterDB(ctrl)
	mockLetter.EXPECT().GetSendedLetters("dellvin.black@gmail.com").Return(nil, nil)
	uc:=LetterUseCase.UseCase{Db: mockLetter}
	uc.GetSendedLetters("dellvin.black@gmail.com")
}
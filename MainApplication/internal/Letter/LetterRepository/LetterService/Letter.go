package LetterService

import (
	"MainApplication/internal/Letter/LetterModel"
	"MainApplication/internal/Letter/LetterRepository"
	"MainApplication/internal/pkg/convert"
	"MainApplication/proto/MailService"
	"context"
)

type LetterServiceManager struct {
	lsClient letterService.LetterServiceClient
}

func New(client letterService.LetterServiceClient) LetterRepository.LetterDB {
	return LetterServiceManager{lsClient: client}
}

func (lsManager LetterServiceManager) WatchLetter(lid uint64) (error, LetterModel.Letter) {
	ctx := context.Background()
	lr, err := lsManager.lsClient.WatchedLetter(ctx, &letterService.Lid{Lid: lid})
	if lr.Letter == nil {
		return LetterRepository.WatchLetterError, LetterModel.Letter{}
	}
	return err, convert.ProtoToModel(lr.Letter)
}

func (lsManager LetterServiceManager) SaveMail(letter LetterModel.Letter) error {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.SaveLetter(ctx, convert.ModelToProto(letter))
	if resp.Ok == false {
		return LetterRepository.SaveLetterError
	}
	return nil
}
func (lsManager LetterServiceManager) GetReceivedLetters(email string) (error, []LetterModel.Letter) {
	ctx := context.Background()
	resp, err := lsManager.lsClient.GetLettersRecv(ctx, &letterService.Email{Email: email})
	if err != nil {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, convert.ProtoToModelList(resp.Letter)
}
func (lsManager LetterServiceManager) GetSendedLetters(email string) (error, []LetterModel.Letter) {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.GetLettersSend(ctx, &letterService.Email{Email: email})
	if resp.Result.Ok == false {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, convert.ProtoToModelList(resp.Letter)
}

func (lsManager LetterServiceManager) GetReceivedLettersDir(dir uint64) (error, []LetterModel.Letter) {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.GetLettersByDirRecv(ctx, &letterService.DirName{DirName: dir})
	if resp.Result.Ok == false {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, convert.ProtoToModelList(resp.Letter)
}

func (lsManager LetterServiceManager) GetSendedLettersDir(dir uint64) (error, []LetterModel.Letter) {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.GetLettersByDirSend(ctx, &letterService.DirName{DirName: dir})
	if resp.Result.Ok == false {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, convert.ProtoToModelList(resp.Letter)
}

func (lsManager LetterServiceManager) DeleteLetter(lid uint64) error{
	ctx := context.Background()
	resp, _:=lsManager.lsClient.RemoveLetter(ctx, &letterService.Lid{Lid: lid})
	if resp.Ok==false{
		return LetterRepository.DeleteLetterError
	}
	return nil
}
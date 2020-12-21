package LetterService

import (
	"Mailer/MainApplication/internal/Letter/LetterModel"
	"Mailer/MainApplication/internal/Letter/LetterRepository"
	"Mailer/MainApplication/internal/pkg/convert"
	"Mailer/MainApplication/proto/MailService"
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
func (lsManager LetterServiceManager) GetReceivedLetters(email string, limit uint64, offset uint64) (error, []LetterModel.Letter) {
	ctx := context.Background()
	resp, err := lsManager.lsClient.GetLettersRecv(ctx, &letterService.Email{
		Email: email,
		Limit: limit,
		Offset: offset,
	})
	if err != nil {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, convert.ProtoToModelList(resp.Letter)
}

func (lsManager LetterServiceManager) GetSendedLetters(email string, limit uint64, offset uint64) (error, []LetterModel.Letter) {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.GetLettersSend(ctx,&letterService.Email{
		Email: email,
		Limit: limit,
		Offset: offset,
	})
	if resp.Result.Ok == false {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, convert.ProtoToModelList(resp.Letter)
}

func (lsManager LetterServiceManager) GetReceivedLettersDir(dir uint64, limit uint64, offset uint64) (error, []LetterModel.Letter) {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.GetLettersByDirRecv(ctx, &letterService.DirName{
		DirName: dir,
		Limit: limit,
		Offset: offset,
		})
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

func (lsManager LetterServiceManager) FindSimilar(sim string, email string) string{
	ctx := context.Background()
	resp, _:=lsManager.lsClient.FindSimilar(ctx, &letterService.Similar{Sim: sim, Email: email})
	return resp.Res
}

func (lsManager LetterServiceManager) GetLetterBy(what string, val string) (error, []LetterModel.Letter){
	ctx := context.Background()
	resp, _:=lsManager.lsClient.GetLetterBy(ctx, &letterService.GetBy{What: what, Value: val})
	if resp.Result.Ok == false {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, convert.ProtoToModelList(resp.Letter)
}
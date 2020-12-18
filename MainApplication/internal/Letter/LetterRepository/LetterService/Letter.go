package LetterService

import (
	msProto "Mailer/MailService/proto"
	"Mailer/MainApplication/internal/Letter/LetterModel"
	"Mailer/MainApplication/internal/Letter/LetterRepository"
	"Mailer/MainApplication/internal/pkg/convert"
	"context"
)

type LetterServiceManager struct {
	lsClient msProto.LetterServiceClient
}

func New(client msProto.LetterServiceClient) LetterRepository.LetterDB {
	return LetterServiceManager{lsClient: client}
}

func (lsManager LetterServiceManager) WatchLetter(lid uint64) (error, LetterModel.Letter) {
	ctx := context.Background()
	lr, err := lsManager.lsClient.WatchedLetter(ctx, &msProto.Lid{Lid: lid})
	if lr.Letter == nil {
		return LetterRepository.WatchLetterError, LetterModel.Letter{}
	}
	return err, convert.ProtoToModel(lr.Letter)
}

func (lsManager LetterServiceManager) SaveMail(letter LetterModel.Letter) error {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.SaveLetter(ctx, convert.ModelToProto(letter))
	if !resp.Ok {
		return LetterRepository.SaveLetterError
	}
	return nil
}
func (lsManager LetterServiceManager) GetReceivedLetters(email string, limit uint64, offset uint64) (error, []LetterModel.Letter) {
	ctx := context.Background()
	resp, err := lsManager.lsClient.GetLettersRecv(ctx, &msProto.Email{
		Email:  email,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, convert.ProtoToModelList(resp.Letter)
}
func (lsManager LetterServiceManager) GetSendedLetters(email string) (error, []LetterModel.Letter) {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.GetLettersSend(ctx, &msProto.Email{Email: email})
	if !resp.Result.Ok {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, convert.ProtoToModelList(resp.Letter)
}

func (lsManager LetterServiceManager) GetReceivedLettersDir(dir uint64) (error, []LetterModel.Letter) {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.GetLettersByDirRecv(ctx, &msProto.DirName{DirName: dir})
	if !resp.Result.Ok {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, convert.ProtoToModelList(resp.Letter)
}

func (lsManager LetterServiceManager) GetSendedLettersDir(dir uint64) (error, []LetterModel.Letter) {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.GetLettersByDirSend(ctx, &msProto.DirName{DirName: dir})
	if !resp.Result.Ok {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, convert.ProtoToModelList(resp.Letter)
}

func (lsManager LetterServiceManager) DeleteLetter(lid uint64) error {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.RemoveLetter(ctx, &msProto.Lid{Lid: lid})
	if !resp.Ok {
		return LetterRepository.DeleteLetterError
	}
	return nil
}

func (lsManager LetterServiceManager) FindSimilar(sim string) string {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.FindSimilar(ctx, &msProto.Similar{Sim: sim})
	return resp.Res
}

func (lsManager LetterServiceManager) GetLetterBy(what string, val string) (error, []LetterModel.Letter) {
	ctx := context.Background()
	resp, _ := lsManager.lsClient.GetLetterBy(ctx, &msProto.GetBy{What: what, Value: val})
	if !resp.Result.Ok {
		return LetterRepository.ReceivedLetterError, nil
	}
	return nil, convert.ProtoToModelList(resp.Letter)
}

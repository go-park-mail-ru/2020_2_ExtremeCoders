package UseCase

import (
	"Mailer/MailService/internal/Model"
	"Mailer/MailService/internal/Repository"
	"Mailer/MailService/pkg/Domain"
	"Mailer/config"
)

//go:generate mockgen -source=./LetterRepository.go -destination=../../test/mock_LetterUseCase/RepositoryMock.go

type UseCase struct {
	re Repository.LetterDB
}

type Interface interface {
	GetLettersSendDir(dir uint64) (error, []Model.Letter)
	GetLettersRecvDir(dir uint64, limit uint64, offset uint64) (error, []Model.Letter)
	GetLettersSend(email string, limit uint64, offset uint64) (error, []Model.Letter)
	GetLettersRecv(email string, limit uint64, offset uint64) (error, []Model.Letter)
	SaveLetter(letter Model.Letter) error
	WatchLetter(lid uint64) (error, Model.Letter)

	AddLetterToDir(uint64, uint64, bool) error
	RemoveLetterFromDir(uint64, uint64, bool) error
	RemoveDir(uint64, bool) error
	RemoveLetter(uint64) error
	FindSimilar(similar string, email string) SearchResult
	GetLetterBy(what string, val string, email string) (error, []Model.Letter)

	SetLetterInSpam(lid uint64) error
	SetLetterInBox(lid uint64) error
}


func New(repo Repository.LetterDB) Interface {
	return UseCase{re: repo}
}

func (uc UseCase) GetLettersRecvDir(dir uint64, limit uint64, offset uint64) (error, []Model.Letter) {
	err, letters := uc.re.GetLettersRecvDir(dir, limit, offset)
	return err, letters
}

func (uc UseCase) GetLettersSendDir(dir uint64) (error, []Model.Letter) {
	err, letters := uc.re.GetLettersSentDir(dir)
	return err, letters
}

func (uc UseCase) SaveLetter(letter Model.Letter) error {
	var err error
	if Domain.GetMailDomain(letter.Receiver)!=config.MailDomain{
		err= uc.re.SendOnAnotherDomain(letter)
		if err!=nil{
			letter.Receiver+=" -> UNKNOWN ADDRESS"
		}
	}
	_=uc.re.SaveMail(letter)
	return err
}

func (uc UseCase) WatchLetter(lid uint64) (error, Model.Letter) {
	return uc.re.SetLetterWatched(lid)
}

func (uc UseCase) GetLettersSend(email string, limit uint64, offset uint64) (error, []Model.Letter) {
	err, letters := uc.re.GetLettersSent(email, limit, offset)
	return err, letters
}

func (uc UseCase) GetLettersRecv(email string, limit uint64, offset uint64) (error, []Model.Letter) {
	err, letters := uc.re.GetLettersRecv(email, limit, offset)
	return err, letters
}

func (uc UseCase) AddLetterToDir(lid uint64, did uint64, flag bool) error {
	return uc.re.AddLetterToDir(lid, did, flag)
}
func (uc UseCase) RemoveLetterFromDir(lid uint64, did uint64, flag bool) error {
	return uc.re.RemoveLetterFromDir(lid, did, flag)
}
func (uc UseCase) RemoveDir(did uint64, flag bool) error {
	return uc.re.RemoveDir(did, flag)
}

func (uc UseCase) RemoveLetter(lid uint64) error{
	return uc.re.RemoveLetter(lid)
}

func (uc UseCase) FindSimilar(similar string, email string) SearchResult {
	res:= SearchResult{}
	res.SimilarTo=similar
	recv, err:=uc.re.FindReceiver(similar, email)
	if err==nil{
		res.Receivers=recv
	}
	send, err:=uc.re.FindSender(similar, email)
	if err==nil{
		res.Senders=send
	}
	theme, err:=uc.re.FindTheme(similar, email)
	if err==nil{
		res.Themes=theme
	}
	text, err:=uc.re.FindText(similar, email)
	if err==nil{
		res.Texts=text
	}
	return res
}

func (uc UseCase) GetLetterBy(what string, val string, email string) (error, []Model.Letter){
	switch what {
	case "sender":
		return uc.re.GetLetterBySender(val, email)
	case "receiver":
		return uc.re.GetLetterByReceiver(val, email)
	case "theme":
		return uc.re.GetLetterByTheme(val, email)
	case "text":
		return uc.re.GetLetterByText(val, email)
	case "spam":
		return uc.re.GetSpam(email)
	case "box":
		return uc.re.GetBox(email)
	}
	return Repository.GetLetterByError, nil
}

func (uc UseCase) SetLetterInSpam(lid uint64) error{
	return uc.re.SetItSpam(lid)
}

func (uc UseCase) SetLetterInBox(lid uint64) error{
	return uc.re.SetItBox(lid)
}

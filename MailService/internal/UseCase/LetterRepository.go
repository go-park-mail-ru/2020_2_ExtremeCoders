package UseCase

import (
	"MailService/internal/Model"
	"MailService/internal/Repository"
)

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
	GetLetterBy(what string, val string) (error, []Model.Letter)
}

//go:generate mockgen -source=./LetterRepository.go -destination=./RepositoryMock.go

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
	err := uc.re.SaveMail(letter)
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

func (uc UseCase) GetLetterBy(what string, val string) (error, []Model.Letter){
	switch what {
	case "sender":
		return uc.re.GetLetterBySender(val)
	case "receiver":
		return uc.re.GetLetterByReceiver(val)
	case "theme":
		return uc.re.GetLetterByTheme(val)
	case "text":
		return uc.re.GetLetterByText(val)
	}
	return Repository.GetLetterByError, nil
}
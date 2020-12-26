package LetterPostgres

import (
	"Mailer/MailService/internal/Model"
	"Mailer/MailService/internal/Repository"
	"Mailer/MailService/pkg/convert"
	smtp "Mailer/SmtpService/proto"
	"context"
	crypto "crypto/rand"
	"fmt"
	"github.com/pkg/errors"
	pgwrapper "gitlab.com/slax0rr/go-pg-wrapper"
	"google.golang.org/grpc"
	"math/big"
	"sort"
)

type dataBase struct {
	DB pgwrapper.DB
}

func New(db pgwrapper.DB) Repository.LetterDB {
	return dataBase{DB: db}
}

func (dbInfo dataBase) SaveMail(letter Model.Letter) error {
	if letter.Receiver==letter.Sender{
		letter.IsWatched=true
	}
	_, err := dbInfo.DB.Model(&letter).Insert()
	if err != nil {
		return Repository.SaveLetterError
	}
	return nil
}

func (dbInfo dataBase) GenerateLID() uint64 {
	for {
		lid, _ := crypto.Int(crypto.Reader, big.NewInt(4294967295))
		user := Model.Letter{Id: lid.Uint64()}
		exist := dbInfo.DB.Model(user).Where("id=?", lid.Int64()).Select()
		if exist != nil {
			return lid.Uint64()
		}
	}
}

func (dbInfo dataBase) GetLettersByFolder(did uint64) (error, []Model.Letter) {
	letters := []Model.Letter{}
	exist := dbInfo.DB.Model(&letters).Where("directory_recv=? or directory_send=?", did, did).
		Select()
	if exist != nil {
		return Repository.SentLetterError, nil
	}
	return nil, letters
}

func (dbInfo dataBase) SetLetterWatched(lid uint64) (error, Model.Letter) {
	err, letter := dbInfo.GetLetterByLid(lid)
	if err != nil {
		return Repository.GetByLidError, letter
	}
	fmt.Println("SET LETTER WATCH", letter.IsWatched)
	letter.IsWatched = !letter.IsWatched
	_, err = dbInfo.DB.Model(&letter).Column("is_watched").Where("id=?", lid).Update()
	if err != nil {
		return Repository.SetLetterWatchedError, letter
	}
	return nil, letter
}

func (dbInfo dataBase) GetLetterByLid(lid uint64) (error, Model.Letter) {
	var letter Model.Letter
	exist := dbInfo.DB.Model(&letter).Where("id=?", lid).Select()
	if exist != nil {
		return Repository.SentLetterError, letter
	}
	return nil, letter
}

func (dbInfo dataBase) GetLettersRecvDir(Did uint64, limit uint64, offset uint64) (error, []Model.Letter) {
	letters := []Model.Letter{}
	exist := dbInfo.DB.Model(&letters).Where("directory_recv=?", Did).Order("date_time DESC").
		Limit(int(limit)).Offset(int(offset)).Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}
	sort.Slice(letters, func(i, j int) (less bool) {
		return letters[i].DateTime > letters[j].DateTime
	})
	data:= []Model.Letter{}
	for _, let:=range letters{
		if !let.Box && !let.Spam{
			data=append(data, let)
		}
	}
	return nil, data
}

func (dbInfo dataBase) GetLettersSentDir(Did uint64, limit uint64, offset uint64) (error, []Model.Letter) {
	letters := []Model.Letter{}
	exist := dbInfo.DB.Model(&letters).Where("directory_send=?", Did).Order("date_time DESC").
		Limit(int(limit)).Offset(int(offset)).Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}
	data:= []Model.Letter{}
	for _, let:=range letters{
		if !let.Box && !let.Spam{
			data=append(data, let)
		}
	}
	return nil, data
}

func (dbInfo dataBase) GetLettersRecv(email string, limit uint64, offset uint64) (error, []Model.Letter) {
	letters := []Model.Letter{}
	exist := dbInfo.DB.Model(&letters).Where("receiver=?", email).Order("date_time DESC").
		Limit(int(limit)).Offset(int(offset)).Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}

	sort.Slice(letters, func(i, j int) (less bool) {
		return letters[i].DateTime > letters[j].DateTime
	})
	data:= []Model.Letter{}
	for _, let:=range letters{
		if !let.Box && !let.Spam{
			data=append(data, let)
		}
	}
	return nil, data
}

func (dbInfo dataBase) GetLettersSent(email string, limit uint64, offset uint64) (error, []Model.Letter) {
	letters := []Model.Letter{}
	exist := dbInfo.DB.Model(&letters).Where("sender=?", email).Order("date_time DESC").
		Limit(int(limit)).Offset(int(offset)).Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}
	sort.Slice(letters, func(i, j int) (less bool) {
		return letters[i].DateTime > letters[j].DateTime
	})
	return nil, letters
}

func (dbInfo dataBase) AddLetterToDir(lid uint64, did uint64, flag bool) error {
	err, letter := dbInfo.GetLetterByLid(lid)
	if err != nil {
		return err
	}
	letter.Box=false
	letter.Spam=false
	if flag {
		letter.DirectoryRecv = did
		_, err = dbInfo.DB.Model(&letter).Column("directory_recv").Where("id=?", lid).
			Update()
		if err != nil {
			return err
		}
	} else {
		letter.DirectorySend = did
		_, err = dbInfo.DB.Model(&letter).Column("directory_send").Where("id=?", lid).
			Update()
		if err != nil {
			return err
		}
	}
	return nil
}
func (dbInfo dataBase) RemoveLetterFromDir(lid uint64, did uint64, flag bool) error {
	err, letter := dbInfo.GetLetterByLid(lid)
	if err != nil {
		return err
	}
	if flag {
		letter.DirectoryRecv = 0
		_, err = dbInfo.DB.Model(&letter).Column("directory_recv").
			Where("id=?", lid).Update()
		if err != nil {
			return err
		}
	} else {
		letter.DirectorySend = 0
		_, err = dbInfo.DB.Model(&letter).Column("directory_send").
			Where("id=?", lid).Update()
		if err != nil {
			return err
		}
	}

	return nil
}

func (dbInfo dataBase) RemoveDir(did uint64, flag bool) error {
	err, letters := dbInfo.GetLettersByFolder(did)
	if err != nil {
		return err
	}
	for _, letter := range letters {
		if flag {
			letter.DirectoryRecv = 0
			_, err = dbInfo.DB.Model(&letter).Column("directory_recv").
				Where("id=?", letter.Id).Update()
		} else {
			letter.DirectorySend = 0
			_, err = dbInfo.DB.Model(&letter).Column("directory_send").
				Where("id=?", letter.Id).Update()
		}
	}
	return err
}

func (dbInfo dataBase) RemoveLetter(lid uint64) error {
	err, letter := dbInfo.GetLetterByLid(lid)
	if err != nil {
		return err
	}
	_, err = dbInfo.DB.Model(&letter).Where("id=?", lid).Delete()
	if err != nil {
		return Repository.DeleteLetterError
	}
	return nil
}

func (dbInfo dataBase) FindSender(theme string, email string) ([]string, error) {
	var letter []Model.Letter
	_, err := dbInfo.DB.Query(&letter, "SELECT * FROM letters WHERE sender LIKE '%' || ? || '%'", theme)
	if err != nil {
		return nil, err
	}
	var data []string
	for _, let := range letter {
		pos := sort.SearchStrings(data, let.Sender)
		if pos == len(data) {
			data = append(data, let.Sender)
		}
	}
	return data, nil
}

func (dbInfo dataBase) FindReceiver(theme string, email string) ([]string, error) {
	var letter []Model.Letter
	_, err := dbInfo.DB.Query(&letter, "SELECT * FROM letters WHERE receiver LIKE '%' || ? || '%'", theme)
	if err != nil {
		return nil, err
	}
	var data []string
	for _, let := range letter {
		pos := sort.SearchStrings(data, let.Receiver)
		if pos == len(data) {
			data = append(data, let.Receiver)
		}
	}
	return data, nil
}

func (dbInfo dataBase) FindTheme(theme string, email string) ([]string, error) {
	var letter []Model.Letter
	_, err := dbInfo.DB.Query(&letter, "SELECT * FROM letters WHERE theme LIKE '%' || ? || '%'", theme)
	if err != nil {
		return nil, err
	}

	var data []string
	for _, let := range letter {
		if let.Receiver == email || let.Sender == email {
			pos := sort.SearchStrings(data, let.Theme)
			if pos == len(data) {
				data = append(data, let.Theme)
			}
		}
	}
	return data, nil
}

func (dbInfo dataBase) FindText(text string, email string) ([]string, error) {
	var letter []Model.Letter
	_, err := dbInfo.DB.Query(&letter, "SELECT DISTINCT * FROM letters WHERE text LIKE '%' || ? || '%'", text)
	if err != nil {
		return nil, err
	}

	var data []string
	for _, let := range letter {
		if let.Receiver == email || let.Sender == email {
			pos := sort.SearchStrings(data, let.Theme)
			if pos == len(data) {
				data = append(data, let.Text)
			}
		}
	}
	return data, nil
}

func (dbInfo dataBase) GetLetterByTheme(val string, email string, limit uint64, offset uint64) (error, []Model.Letter) {
	var letters []Model.Letter
	err := dbInfo.DB.Model(&letters).Where("theme=?", val).Order("date_time DESC").
		Limit(int(limit)).Offset(int(offset)).Select()
	if err != nil {
		return Repository.GetLetterByError, nil
	}
	var data []Model.Letter
	for _, let := range letters {
		if let.Receiver == email || let.Sender == email {
			data = append(data, let)
		}
	}
	return nil, data
}

func (dbInfo dataBase) GetLetterByText(val string, email string, limit uint64, offset uint64) (error, []Model.Letter) {
	var letters []Model.Letter
	err := dbInfo.DB.Model(&letters).Where("text=?", val).Order("date_time DESC").
		Limit(int(limit)).Offset(int(offset)).Select()
	if err != nil {
		return Repository.GetLetterByError, nil
	}
	var data []Model.Letter
	for _, let := range letters {
		if let.Receiver == email || let.Sender == email {
			data = append(data, let)
		}
	}
	return nil, data
}

func (dbInfo dataBase) GetLetterBySender(val string, email string, limit uint64, offset uint64) (error, []Model.Letter) {
	var letters []Model.Letter
	err := dbInfo.DB.Model(&letters).Where("sender=?", val).Order("date_time DESC").
		Limit(int(limit)).Offset(int(offset)).Select()
	if err != nil {
		return Repository.GetLetterByError, nil
	}
	var data []Model.Letter
	for _, let := range letters {
		if let.Receiver == email || let.Sender == email {
			data = append(data, let)
		}
	}
	return nil, data
}

func (dbInfo dataBase) GetLetterByReceiver(val string, email string, limit uint64, offset uint64) (error, []Model.Letter) {
	var letters []Model.Letter
	err := dbInfo.DB.Model(&letters).Where("receiver=?", val).Order("date_time DESC").
		Limit(int(limit)).Offset(int(offset)).Select()
	if err != nil {
		return Repository.GetLetterByError, nil
	}
	var data []Model.Letter
	for _, let := range letters {
		if let.Receiver == email || let.Sender == email {
			data = append(data, let)
		}
	}
	return nil, data
}


func (dbInfo dataBase) GetSpam(email string, limit uint64, offset uint64) (error, []Model.Letter) {
	var letters []Model.Letter
	err := dbInfo.DB.Model(&letters).Where("spam=?", true).Order("date_time DESC").
		Limit(int(limit)).Offset(int(offset)).Select()
	if err != nil {
		return Repository.GetLetterByError, nil
	}
	var data []Model.Letter
	for _, let := range letters {
		if (let.Receiver == email || let.Sender == email) && !let.Box {
			data = append(data, let)
		}
	}
	return nil, data
}

func (dbInfo dataBase) GetBox(email string, limit uint64, offset uint64) (error, []Model.Letter) {
	var letters []Model.Letter
	err := dbInfo.DB.Model(&letters).Where("box=?", true).Order("date_time DESC").
		Limit(int(limit)).Offset(int(offset)).Select()
	if err != nil {
		return Repository.GetLetterByError, nil
	}
	var data []Model.Letter
	for _, let := range letters {
		if let.Receiver == email || let.Sender == email {
			data = append(data, let)
		}
	}
	return nil, data
}

func (dbInfo dataBase)SetItSpam(lid uint64) error{
	err, letter:= dbInfo.GetLetterByLid(lid)
	if letter.Spam{
		letter.Spam=false
	}
	if err!=nil{
		return Repository.GetByLidError
	}
	letter.Spam=true
	_, err=dbInfo.DB.Model(&letter).Column("spam").Where("id=?", lid).Update()
	if err!=nil{
		return Repository.SetSpamError
	}
	return nil
}

func (dbInfo dataBase)SetItBox(lid uint64) error{
	err, letter:= dbInfo.GetLetterByLid(lid)
	if letter.Box{
		letter.Box=false
	} else{
		letter.Spam=false
	}
	if err!=nil{
		return Repository.GetByLidError
	}
	letter.Box=true
	_, err=dbInfo.DB.Model(&letter).Column("box").Where("id=?", lid).Update()
	if err!=nil{
		return Repository.SetBoxError
	}
	return nil
}

func (dbInfo dataBase)SendOnAnotherDomain(letter Model.Letter) error{
	grcpMailService, err := grpc.Dial(
		"147.78.67.180:8080",
		grpc.WithInsecure(),
	)
	if err!=nil{
		fmt.Println("HIU: ", err.Error())
		return err
	}
	defer grcpMailService.Close()
	mailManager :=smtp.NewLetterServiceClient(grcpMailService)
	ctx:=context.Background()
	fmt.Println("connect to smpt serv")
	resp, _:=mailManager.SendLetter(ctx, convert.ModelToSmtp(letter))
	if resp!=nil && !resp.Ok{
		var err error
		return errors.Wrapf(err, resp.Description)
	}
	return nil
}
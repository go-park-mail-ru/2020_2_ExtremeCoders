package LetterPostgres

import (
	"Mailer/MailService/internal/Model"
	"Mailer/MailService/internal/Repository"
	crypto "crypto/rand"
	pgwrapper "gitlab.com/slax0rr/go-pg-wrapper"
	"sort"

	"math/big"
)

type dataBase struct {
	DB pgwrapper.DB
}

func New(db pgwrapper.DB) Repository.LetterDB {
	return dataBase{DB: db}
}

func (dbInfo dataBase) SaveMail(letter Model.Letter) error {
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
	letter.IsWatched = true
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
	exist := dbInfo.DB.Model(&letters).Where("directory_recv=?", Did).
		Limit(int(limit)).Offset(int(offset)).Order("date_time DESC").Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}
	return nil, letters
}

func (dbInfo dataBase) GetLettersSentDir(Did uint64) (error, []Model.Letter) {
	letters := []Model.Letter{}
	exist := dbInfo.DB.Model(&letters).Where("directory_send=?", Did).Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}
	return nil, letters
}

func (dbInfo dataBase) GetLettersRecv(email string, limit uint64, offset uint64) (error, []Model.Letter) {
	letters := []Model.Letter{}
	exist := dbInfo.DB.Model(&letters).Where("receiver=?", email).
		Limit(int(limit)).Offset(int(offset)).Order("date_time DESC").Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}
	return nil, letters
}

func (dbInfo dataBase) GetLettersSent(email string, limit uint64, offset uint64) (error, []Model.Letter) {
	letters := []Model.Letter{}
	exist := dbInfo.DB.Model(&letters).Where("sender=?", email).
		Limit(int(limit)).Offset(int(offset)).Order("date_time DESC").Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}
	return nil, letters
}

func (dbInfo dataBase) AddLetterToDir(lid uint64, did uint64, flag bool) error {
	err, letter := dbInfo.GetLetterByLid(lid)
	if err != nil {
		return err
	}
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

func (dbInfo dataBase) RemoveLetter(lid uint64) error{
	err, letter:=dbInfo.GetLetterByLid(lid)
	if err!=nil{
		return err
	}
	_, err=dbInfo.DB.Model(&letter).Where("id=?", lid).Delete()
	if err!=nil{
		return Repository.DeleteLetterError
	}
	return nil
}

func (dbInfo dataBase)FindSender(theme string, email string) ([]string, error){
	var letter []Model.Letter
	_, err := dbInfo.DB.Query(&letter, "SELECT * FROM letters WHERE sender LIKE '%' || ? || '%'", theme)
	if err!=nil{
		return nil, err
	}
	var data []string
	for _, let:=range letter{
		pos:=sort.SearchStrings(data, let.Sender)
		if pos== len(data){
			data=append(data, let.Sender)
		}
	}
	return data, nil
}

func (dbInfo dataBase)FindReceiver(theme string, email string) ([]string, error){
	var letter []Model.Letter
	_, err := dbInfo.DB.Query(&letter, "SELECT * FROM letters WHERE receiver LIKE '%' || ? || '%'", theme)
	if err!=nil{
		return nil, err
	}
	var data []string
	for _, let:=range letter{
		pos:=sort.SearchStrings(data, let.Receiver)
		if pos== len(data){
			data=append(data, let.Receiver)
		}
	}
	return data, nil
}

func (dbInfo dataBase)FindTheme(theme string, email string) ([]string, error){
	var letter []Model.Letter
	_, err := dbInfo.DB.Query(&letter, "SELECT * FROM letters WHERE theme LIKE '%' || ? || '%'", theme)
	if err!=nil{
		return nil, err
	}

	var data []string
	for _, let:=range letter{
		if let.Receiver==email || let.Sender==email{
			pos:=sort.SearchStrings(data, let.Theme)
			if pos== len(data){
				data=append(data, let.Theme)
			}
		}
	}
	return data, nil
}

func (dbInfo dataBase)FindText(text string, email string) ([]string, error){
	var letter []Model.Letter
	_, err := dbInfo.DB.Query(&letter, "SELECT DISTINCT * FROM letters WHERE text LIKE '%' || ? || '%'", text)
	if err!=nil{
		return nil, err
	}

	var data []string
	for _, let:=range letter{
		if let.Receiver==email || let.Sender==email{
			pos:=sort.SearchStrings(data, let.Theme)
			if pos== len(data){
				data=append(data, let.Text)
			}
		}
	}
	return data, nil
}

func (dbInfo dataBase) GetLetterByTheme(val string) (error, []Model.Letter){
	var letters []Model.Letter
	_, err := dbInfo.DB.Query(&letters, "SELECT * FROM letters WHERE theme = ?", val)

	if err!=nil{
		return Repository.GetLetterByError, nil
	}
	return nil, letters
}

func (dbInfo dataBase) GetLetterByText(val string) (error, []Model.Letter){
	var letters []Model.Letter
	_, err := dbInfo.DB.Query(&letters, "SELECT * FROM letters WHERE text = ?", val)

	if err!=nil{
		return Repository.GetLetterByError, nil
	}
	return nil, letters
}

func (dbInfo dataBase) GetLetterBySender(val string) (error, []Model.Letter){
	var letters []Model.Letter
	_, err := dbInfo.DB.Query(&letters, "SELECT * FROM letters WHERE sender = ?", val)

	if err!=nil{
		return Repository.GetLetterByError, nil
	}
	return nil, letters
}

func (dbInfo dataBase) GetLetterByReceiver(val string) (error, []Model.Letter){
	var letters []Model.Letter
	_, err := dbInfo.DB.Query(&letters, "SELECT * FROM letters WHERE receiver = ?", val)

	if err!=nil{
		return Repository.GetLetterByError, nil
	}
	return nil, letters
}
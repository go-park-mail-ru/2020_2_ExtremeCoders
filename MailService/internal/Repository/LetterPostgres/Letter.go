package LetterPostgres

import (
	"MailService/internal/Model"
	"MailService/internal/Repository"
	crypto "crypto/rand"
	pgwrapper "gitlab.com/slax0rr/go-pg-wrapper"
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
	var letters []Model.Letter
	exist := dbInfo.DB.Model(&letters).Where("directory_recv=? or directory_send=?", did, did).Select()
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

func (dbInfo dataBase) GetLettersRecvDir(Did uint64) (error, []Model.Letter) {
	var letters []Model.Letter
	exist := dbInfo.DB.Model(&letters).Where("directory_recv=?", Did).Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}
	return nil, letters
}

func (dbInfo dataBase) GetLettersSentDir(Did uint64) (error, []Model.Letter) {
	var letters []Model.Letter
	exist := dbInfo.DB.Model(&letters).Where("directory_send=?", Did).Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}
	return nil, letters
}

func (dbInfo dataBase) GetLettersRecv(email string) (error, []Model.Letter) {
	var letters []Model.Letter
	exist := dbInfo.DB.Model(&letters).Where("receiver=?", email).Select()
	if exist != nil {
		return Repository.SentLetterError, letters
	}
	return nil, letters
}

func (dbInfo dataBase) GetLettersSent(email string) (error, []Model.Letter) {
	var letters []Model.Letter
	exist := dbInfo.DB.Model(&letters).Where("sender=?", email).Select()
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
		_, err = dbInfo.DB.Model(&letter).Column("directory_recv").Where("id=?", lid).Update()
		if err != nil {
			return err
		}
	} else {
		letter.DirectorySend = did
		_, err = dbInfo.DB.Model(&letter).Column("directory_send").Where("id=?", lid).Update()
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
		_, err = dbInfo.DB.Model(&letter).Column("directory_recv").Where("id=?", lid).Update()
		if err != nil {
			return err
		}
	} else {
		letter.DirectorySend = 0
		_, err = dbInfo.DB.Model(&letter).Column("directory_send").Where("id=?", lid).Update()
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
			_, err = dbInfo.DB.Model(&letter).Column("directory_recv").Where("id=?", letter.Id).Update()
		} else {
			letter.DirectorySend = 0
			_, err = dbInfo.DB.Model(&letter).Column("directory_send").Where("id=?", letter.Id).Update()
		}
	}
	return err
}

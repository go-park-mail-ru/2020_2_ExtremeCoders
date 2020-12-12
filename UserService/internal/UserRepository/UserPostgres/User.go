package UserPostgres

//go:generate mockgen -destination=../mocks/mock_User.go -package=mocks -source=./DataBaseRequests.go

import (
	"Mailer/config"
	"Mailer/UserService/internal/UserModel"
	"Mailer/UserService/internal/UserRepository"
	crypto "crypto/rand"
	"fmt"
	"github.com/go-pg/pg/v9"
	pgwrapper "gitlab.com/slax0rr/go-pg-wrapper"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)

type dataBase struct {
	DB pgwrapper.DB
}

func New(db pgwrapper.DB) UserRepository.UserDB {
	return dataBase{DB: db}
}

func (dbInfo dataBase) GetFoldersList(uid uint64, kind string) (folders []*UserModel.Folder, err error) {
	var res []*UserModel.Folder
	err = dbInfo.DB.Model(&res).Where("uid=? and type=?", uid, kind).Select()
	if err != nil {
		fmt.Println("GET FOLDERS ERROR", err)
		return nil, err
	}
	return res, nil
}

func (dbInfo dataBase) RemoveFolder(id uint64) error {
	folder := &UserModel.Folder{Id: id}
	_, err := dbInfo.DB.Model(folder).Where("id=?", id).Delete()
	return err
}

func (dbInfo dataBase) RenameFolder(uid uint64, kind string, oldName string, newName string) error {
	fmt.Println("CALL RenameFolder")
	oldFolder := &UserModel.Folder{Uid: uid, Name: oldName}
	err := dbInfo.DB.Model(oldFolder).Where("uid=? and name=? and type=?", uid, oldName, kind).Select()
	if err != nil {
		return UserRepository.RenameFolderError
	}
	folder := oldFolder
	folder.Name = newName
	_, err = dbInfo.DB.Model(folder).Column("name").Where("id=?", oldFolder.Id).Update()
	if err != nil {
		return UserRepository.RenameFolderError
	}
	return nil
}

func (dbInfo dataBase) CreateFolder(name string, kind string, uid uint64) error {

	fmt.Println("CALL GET FOLDER ID", uid, kind, name)
	folder := &UserModel.Folder{
		Uid:  uid,
		Type: kind,
		Name: name,
	}
	exist:=dbInfo.DB.Model(folder).Where("type=? and name=? and uid=?", folder.Type, folder.Name, folder.Uid).Select()
	if exist==nil{
		return UserRepository.CreateFolderError
	}
	_, err := dbInfo.DB.Model(folder).Insert()
	if err != nil {
		fmt.Println("ERR", err)
		return UserRepository.CreateFolderError
	}
	return nil
}

func (dbInfo dataBase) GetFolderId(uid uint64, kind string, name string) (fid uint64, err error) {
	fmt.Println("CALL GET FOLDER ID", uid, kind, name)
	folder := &UserModel.Folder{
		Uid:  uid,
		Type: kind,
		Name: name,
	}
	err = dbInfo.DB.Model(folder).Where("type=? and uid=? and name=?", kind, uid, name).Select()
	if err != nil {
		fmt.Println("ERR", err)
		return 0, UserRepository.GetFolderIdError
	}
	return folder.Id, nil
}

func (dbInfo dataBase) IsEmailExists(email string) error {
	fmt.Println("CALL IS EMAIL EXIST")
	user := &UserModel.User{Email: email}
	err := dbInfo.DB.Model(user).Where("email=?", email).Select()
	if err != pg.ErrNoRows {
		return UserRepository.EmailAlreadyExists
	}
	return nil
}

func (dbInfo dataBase) AddUser(user *UserModel.User) error {
	fmt.Println("CALL AddUser")

	user.Password = string(PasswordBcrypt([]byte(user.Password)))
	_, err := dbInfo.DB.Model(user).Insert()
	if err != nil {
		return UserRepository.CantAddUser
	}
	return nil
}

func (dbInfo dataBase) AddSession(sid string, uid uint64, user *UserModel.User) error {
	fmt.Println("CALL AddSession ", "sid", sid, "uid", uid, "userUid", user.Id)

	session := &UserModel.Session{Id: sid, UserId: int64(uid), User: user}
	_, err := dbInfo.DB.Model(session).Insert()
	if err != nil {
		return UserRepository.CantAddSession
	}
	return nil
}

func (dbInfo dataBase) GenerateSID() (sessionId []rune, err error) {
	fmt.Println("CALL GenerateSID")
	var sid string
	for {
		for i := 0; i < config.SizeSID; i++ {
			safeNum, _ := crypto.Int(crypto.Reader, big.NewInt(int64(len(config.SidRunes))))
			sid += string(config.SidRunes[safeNum.Int64()])
		}
		fmt.Println(sid)
		session := &UserModel.Session{Id: sid}
		exist := dbInfo.DB.Model(session).Where("id=?", sid).Select()
		if exist != nil {
			break
		}
		sid = ""
	}
	return []rune(sid), nil
}

func (dbInfo dataBase) GenerateUID() (uid uint64, err error) {
	fmt.Println("CALL GenerateUID")
	for {
		uid, _ := crypto.Int(crypto.Reader, big.NewInt(4294967295))
		user := UserModel.User{Id: uid.Uint64()}
		exist := dbInfo.DB.Model(user).Where("id=?", uid.Int64()).Select()
		if exist != nil {
			return uid.Uint64(), nil
		}
	}
}

func (dbInfo dataBase) GetUserByEmail(email string) (user *UserModel.User, err error) {
	fmt.Println("CALL GetUserByEmail")
	user = &UserModel.User{Email: email}
	err = dbInfo.DB.Model(user).Where("email=?", email).Select()
	if err != nil {
		return user, UserRepository.CantGetUserByEmail
	}
	return user, nil
}

func (dbInfo dataBase) GetUserByUID(uid uint64) (user *UserModel.User, err error) {
	fmt.Println("CALL GetUserByUID")
	user = &UserModel.User{}
	err = dbInfo.DB.Model(user).Where("id=?", uid).Select()
	if err != nil {
		return user, UserRepository.CantGetUserByUid
	}
	return user, nil
}

func (dbInfo dataBase) IsOkSession(sid string) (uid uint64, err error) {
	fmt.Println("CALL IsOkSession")
	session := &UserModel.Session{Id: sid}
	err = dbInfo.DB.Model(session).Where("id=?", sid).Select()
	if err != nil {
		return 0, UserRepository.InvalidSession
	}
	return uint64(session.UserId), nil
}

func (dbInfo dataBase) UpdateProfile(newUser UserModel.User, email string) error {
	fmt.Println("CALL UpdateProfile")
	oldUser := &UserModel.User{Email: email}
	err := dbInfo.DB.Model(oldUser).Where("email=?", email).Select()
	if err != nil {
		return UserRepository.CantGetUserOnUpdate
	}
	User := oldUser
	User.Name = newUser.Name
	User.Surname = newUser.Surname
	User.Img = newUser.Img
	_, err = dbInfo.DB.Model(User).Column("name", "surname", "img").Where("email=?", email).Update()
	if err != nil {
		return UserRepository.CantUpdateUser
	}
	return nil
}

func (dbInfo dataBase) RemoveSession(sid string) (err error, uid uint64) {
	fmt.Println("CALL RemoveSession")
	session := &UserModel.Session{Id: sid}
	err = dbInfo.DB.Model(session).Where("id=?", sid).Select()
	_, err = dbInfo.DB.Model(session).Where("id=?", sid).Delete()
	if err != nil {
		return UserRepository.RemoveSessionError, 0
	}
	return nil, uint64(session.UserId)
}

func (dbInfo dataBase) GetSessionByUID(uid uint64) (sid string, err error) {
	fmt.Println("CALL GetSessionByUID")
	session := &UserModel.Session{UserId: int64(uid)}
	err = dbInfo.DB.Model(session).Where("user_id=?", uid).Select()
	if err != nil {
		return "", UserRepository.GetSessionError
	}
	return session.Id, nil
}

func PasswordBcrypt(plainPassword []byte) []byte {
	fmt.Println("CALL PasswordBcrypt")
	passBcrypt, _ := bcrypt.GenerateFromPassword(plainPassword, 14)
	return passBcrypt
}

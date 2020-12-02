package UserUseCase

import (
	"UserService/internal/UserModel"
	proto "UserService/proto"
	"errors"
	"fmt"
)
import "UserService/internal/UserRepository"

type Interface interface {
	IsEmailExists(*proto.Email) (*proto.Nothing, error)
	AddSession(*proto.AddSessionMsg) (*proto.Nothing, error)
	AddUser(*proto.User) (*proto.Nothing, error)
	GenerateSID(*proto.Nothing) (*proto.Sid, error)
	GenerateUID(*proto.Nothing) (*proto.Uid, error)
	GetUserByEmail(*proto.Email) (*proto.User, error)
	GetUserByUID(*proto.Uid) (*proto.User, error)
	IsOkSession(*proto.Sid) (*proto.Uid, error)
	UpdateProfile(*proto.UpdateProfileMsg) (*proto.Nothing, error)
	RemoveSession(*proto.Sid) (*proto.Uid, error)
	GetSessionByUID(*proto.Uid) (*proto.Sid, error)
	GetFolderId(*proto.Folder) (*proto.FolderId, error)
	CreateFolder(*proto.Folder) (*proto.Nothing, error)
	RenameFolder(*proto.RenameFolderMsg) (*proto.Nothing, error)
	RemoveFolder(*proto.Folder) (*proto.FolderId, error)
	GetFoldersList(*proto.FolderUidType) (*proto.FolderList, error)
}

var RemoveFolderErr = errors.New("REMOVE FOLDER ERROR")

type UseCase struct {
	db UserRepository.UserDB
}

func New(db UserRepository.UserDB) Interface {
	return UseCase{db: db}
}

func (u UseCase) GetFoldersList(msg *proto.FolderUidType) (*proto.FolderList, error) {
	folders, err := u.db.GetFoldersList(msg.Uid, msg.Type)
	if err != nil {
		return nil, err
	}
	var tmp []*proto.FolderNameType
	for _, value := range folders {
		tmp = append(tmp, &proto.FolderNameType{
			Name: value.Name,
			Type: value.Type,
		})
	}
	return &proto.FolderList{Res: tmp}, err
}

func (u UseCase) RemoveFolder(folder *proto.Folder) (*proto.FolderId, error) {
	id, err := u.db.GetFolderId(folder.Uid, folder.Type, folder.Name)
	if err != nil {
		return nil, RemoveFolderErr
	}
	err = u.db.RemoveFolder(id)
	if err != nil {
		return nil, RemoveFolderErr
	}
	return &proto.FolderId{Id: id}, err
}

func (u UseCase) RenameFolder(msg *proto.RenameFolderMsg) (*proto.Nothing, error) {
	err := u.db.RenameFolder(msg.Uid, msg.Type, msg.OldName, msg.NewName)
	return &proto.Nothing{Dummy: true}, err
}

func (u UseCase) CreateFolder(folder *proto.Folder) (*proto.Nothing, error) {
	err := u.db.CreateFolder(folder.Name, folder.Type, folder.Uid)
	return &proto.Nothing{Dummy: true}, err
}

func (u UseCase) IsEmailExists(email *proto.Email) (*proto.Nothing, error) {
	err := u.db.IsEmailExists(email.Email)
	return &proto.Nothing{Dummy: true}, err
}

func (u UseCase) GetFolderId(msg *proto.Folder) (*proto.FolderId, error) {
	folderId, err := u.db.GetFolderId(msg.Uid, msg.Type, msg.Name)
	if err != nil {
		return nil, err
	}
	return &proto.FolderId{Id: folderId}, nil
}

func (u UseCase) AddSession(msg *proto.AddSessionMsg) (*proto.Nothing, error) {
	fmt.Println("ADD SESSION UC", msg.Sid, msg.User.Uid, msg.User.Name)
	userModel := UserModel.User{
		Id:       msg.User.Uid,
		Name:     msg.User.Name,
		Surname:  msg.User.Surname,
		Email:    msg.User.Email,
		Password: msg.User.Password,
	}
	err := u.db.AddSession(msg.Sid, userModel.Id, &userModel)
	return &proto.Nothing{Dummy: true}, err
}

func (u UseCase) AddUser(user *proto.User) (*proto.Nothing, error) {
	userModel := UserModel.User{
		Id:       user.Uid,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: user.Password,
	}
	err := u.db.AddUser(&userModel)
	return &proto.Nothing{Dummy: true}, err
}

func (u UseCase) GenerateSID(nothing *proto.Nothing) (*proto.Sid, error) {
	sid, err := u.db.GenerateSID()
	if err != nil {
		return nil, err
	}
	return &proto.Sid{Sid: string(sid)}, err
}

func (u UseCase) GenerateUID(nothing *proto.Nothing) (*proto.Uid, error) {
	uid, err := u.db.GenerateUID()
	if err != nil {
		return nil, err
	}
	return &proto.Uid{Uid: uid}, err
}

func (u UseCase) GetUserByEmail(email *proto.Email) (*proto.User, error) {
	user, err := u.db.GetUserByEmail(email.Email)
	if err != nil {
		return nil, err
	}
	res := proto.User{
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Uid:      user.Id,
	}
	return &res, err
}

func (u UseCase) GetUserByUID(uid *proto.Uid) (*proto.User, error) {
	user, err := u.db.GetUserByUID(uid.Uid)
	if err != nil {
		return nil, err
	}
	res := proto.User{
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Uid:      user.Id,
	}
	return &res, err
}

func (u UseCase) IsOkSession(sid *proto.Sid) (*proto.Uid, error) {
	uid, err := u.db.IsOkSession(sid.Sid)
	if err != nil {
		return nil, err
	}
	return &proto.Uid{Uid: uid}, err
}

func (u UseCase) UpdateProfile(msg *proto.UpdateProfileMsg) (*proto.Nothing, error) {
	userModel := UserModel.User{
		Id:       msg.NewUser.Uid,
		Name:     msg.NewUser.Name,
		Surname:  msg.NewUser.Surname,
		Email:    msg.NewUser.Email,
		Password: msg.NewUser.Password,
	}
	err := u.db.UpdateProfile(userModel, msg.Email)
	if err != nil {
		return nil, err
	}
	return &proto.Nothing{Dummy: true}, err
}

func (u UseCase) RemoveSession(sid *proto.Sid) (*proto.Uid, error) {
	err, uid := u.db.RemoveSession(sid.Sid)
	if err != nil {
		return nil, err
	}
	return &proto.Uid{Uid: uid}, err
}

func (u UseCase) GetSessionByUID(uid *proto.Uid) (*proto.Sid, error) {
	sid, err := u.db.GetSessionByUID(uid.Uid)
	if err != nil {
		return nil, err
	}
	return &proto.Sid{Sid: sid}, err
}

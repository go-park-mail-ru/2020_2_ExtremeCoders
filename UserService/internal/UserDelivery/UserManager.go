package UserDelivery

import (
	"context"
	"fmt"
)
import proto "Mailer/UserService/proto"
import "Mailer/UserService/internal/UserUseCase"

type UserManager struct {
	useCase UserUseCase.Interface
}

func New(uc UserUseCase.Interface) proto.UserServiceServer {
	return UserManager{useCase: uc}
}

func (um UserManager) GetFoldersList(ctx context.Context, uid *proto.FolderUidType) (*proto.FolderList, error) {
	fmt.Println("GetFoldersList DELIVERY ")
	folder, err:=um.useCase.GetFoldersList(uid)
	return folder, err
}

func (um UserManager) RemoveFolder(ctx context.Context, folder *proto.Folder) (*proto.FolderId, error) {
	fmt.Println("RemoveFolder DELIVERY ")
	return um.useCase.RemoveFolder(folder)
}

func (um UserManager) RenameFolder(ctx context.Context, msg *proto.RenameFolderMsg) (*proto.Nothing, error) {
	fmt.Println("RenameFolder DELIVERY ")
	return um.useCase.RenameFolder(msg)
}

func (um UserManager) CreateFolder(ctx context.Context, folder *proto.Folder) (*proto.Nothing, error) {
	fmt.Println("CreateFolder DELIVERY ")
	return um.useCase.CreateFolder(folder)
}

func (um UserManager) GetFolderId(ctx context.Context, msg *proto.Folder) (*proto.FolderId, error) {
	fmt.Println("GetFolderId DELIVERY ")
	return um.useCase.GetFolderId(msg)
}

func (um UserManager) IsEmailExists(ctx context.Context, email *proto.Email) (*proto.Nothing, error) {
	fmt.Println("IsEmailExists DELIVERY ")
	return um.useCase.IsEmailExists(email)
}

func (um UserManager) AddSession(ctx context.Context, msg *proto.AddSessionMsg) (*proto.Nothing, error) {
	fmt.Println("ADD SESSION DELIVERY ", msg.Sid, msg.User.Uid, msg.User.Email)
	return um.useCase.AddSession(msg)
}

func (um UserManager) AddUser(ctx context.Context, user *proto.User) (*proto.Nothing, error) {
	fmt.Println("AddUser DELIVERY ")
	return um.useCase.AddUser(user)
}

func (um UserManager) GenerateSID(ctx context.Context, nothing *proto.Nothing) (*proto.Sid, error) {
	fmt.Println("GenerateSID DELIVERY ")
	return um.useCase.GenerateSID(nothing)
}

func (um UserManager) GenerateUID(ctx context.Context, nothing *proto.Nothing) (*proto.Uid, error) {
	fmt.Println("GenerateUID DELIVERY ")
	return um.useCase.GenerateUID(nothing)
}

func (um UserManager) GetUserByEmail(ctx context.Context, email *proto.Email) (*proto.User, error) {
	fmt.Println("GetUserByEmail DELIVERY ")
	return um.useCase.GetUserByEmail(email)
}

func (um UserManager) GetUserByUID(ctx context.Context, uid *proto.Uid) (*proto.User, error) {
	fmt.Println("GetUserByUID DELIVERY ")
	return um.useCase.GetUserByUID(uid)
}

func (um UserManager) IsOkSession(ctx context.Context, sid *proto.Sid) (*proto.Uid, error) {
	fmt.Println("IsOkSession DELIVERY ")
	return um.useCase.IsOkSession(sid)
}

func (um UserManager) UpdateProfile(ctx context.Context, msg *proto.UpdateProfileMsg) (*proto.Nothing, error) {
	fmt.Println("UpdateProfile DELIVERY ")
	return um.useCase.UpdateProfile(msg)
}

func (um UserManager) RemoveSession(ctx context.Context, sid *proto.Sid) (*proto.Uid, error) {
	fmt.Println("RemoveSession DELIVERY ")
	return um.useCase.RemoveSession(sid)
}

func (um UserManager) GetSessionByUID(ctx context.Context, uid *proto.Uid) (*proto.Sid, error) {
	fmt.Println("GetSessionByUID DELIVERY ")
	return um.useCase.GetSessionByUID(uid)
}

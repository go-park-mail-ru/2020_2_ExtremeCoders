package UserDelivery

import (
	"context"
	"fmt"
)
import proto "UserService/proto"
import "UserService/internal/UserUseCase"

type UserManager struct {
	useCase UserUseCase.Interface
}

func New(uc UserUseCase.Interface) proto.UserServiceServer {
	return UserManager{useCase: uc}
}

func (um UserManager) GetFoldersList(ctx context.Context, uid *proto.FolderUidType) (*proto.FolderList, error) {
	folder, err:=um.useCase.GetFoldersList(uid)
	return folder, err
}

func (um UserManager) RemoveFolder(ctx context.Context, folder *proto.Folder) (*proto.FolderId, error) {
	return um.useCase.RemoveFolder(folder)
}

func (um UserManager) RenameFolder(ctx context.Context, msg *proto.RenameFolderMsg) (*proto.Nothing, error) {
	return um.useCase.RenameFolder(msg)
}

func (um UserManager) CreateFolder(ctx context.Context, folder *proto.Folder) (*proto.Nothing, error) {
	return um.useCase.CreateFolder(folder)
}

func (um UserManager) GetFolderId(ctx context.Context, msg *proto.Folder) (*proto.FolderId, error) {
	return um.useCase.GetFolderId(msg)
}

func (um UserManager) IsEmailExists(ctx context.Context, email *proto.Email) (*proto.Nothing, error) {
	return um.useCase.IsEmailExists(email)
}

func (um UserManager) AddSession(ctx context.Context, msg *proto.AddSessionMsg) (*proto.Nothing, error) {
	fmt.Println("ADD SESSION DELIVERY ", msg.Sid, msg.User.Uid, msg.User.Email)
	return um.useCase.AddSession(msg)
}

func (um UserManager) AddUser(ctx context.Context, user *proto.User) (*proto.Nothing, error) {
	return um.useCase.AddUser(user)
}

func (um UserManager) GenerateSID(ctx context.Context, nothing *proto.Nothing) (*proto.Sid, error) {
	return um.useCase.GenerateSID(nothing)
}

func (um UserManager) GenerateUID(ctx context.Context, nothing *proto.Nothing) (*proto.Uid, error) {
	return um.useCase.GenerateUID(nothing)
}

func (um UserManager) GetUserByEmail(ctx context.Context, email *proto.Email) (*proto.User, error) {
	return um.useCase.GetUserByEmail(email)
}

func (um UserManager) GetUserByUID(ctx context.Context, uid *proto.Uid) (*proto.User, error) {
	return um.useCase.GetUserByUID(uid)
}

func (um UserManager) IsOkSession(ctx context.Context, sid *proto.Sid) (*proto.Uid, error) {
	return um.useCase.IsOkSession(sid)
}

func (um UserManager) UpdateProfile(ctx context.Context, msg *proto.UpdateProfileMsg) (*proto.Nothing, error) {
	return um.useCase.UpdateProfile(msg)
}

func (um UserManager) RemoveSession(ctx context.Context, sid *proto.Sid) (*proto.Uid, error) {
	return um.useCase.RemoveSession(sid)
}

func (um UserManager) GetSessionByUID(ctx context.Context, uid *proto.Uid) (*proto.Sid, error) {
	return um.useCase.GetSessionByUID(uid)
}

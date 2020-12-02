package UserDelivery

import (
	"UserService/internal/UserModel"
	mock "UserService/internal/mocks"
	proto "UserService/proto"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testUser = UserModel.User{
		Id:      	1,
		Name: 		"UserName",
		Surname:  	"UserSurname",
		Email:      "UserEmail",
		Password:   "UserPassword",
		Img:     	"",
	}

	testSession = UserModel.Session{
		Id:        	"SessionId",
		UserId: 	1,
		User:		&testUser,
	}

	testFolder = UserModel.Folder{
		Uid:		1,
		Type: 		"FolderType",
		Name:		"FolderName",
	}
)
func TestUserManager_AddSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoUser := &proto.User{
		Uid: 		testUser.Id,
		Name: 		testUser.Name,
		Surname:  	testUser.Surname,
		Email:      testUser.Email,
		Password:   testUser.Password,
	}

	protoSessionMsg := &proto.AddSessionMsg {
		Sid: testSession.Id,
		User: protoUser,
	}

	protoNothing := &proto.Nothing{Dummy: true}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().AddSession(protoSessionMsg).Return(protoNothing, nil)

	output, err := ud.AddSession(ctx, protoSessionMsg)

	assert.Equal(t, protoNothing, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoUser := &proto.User{
		Uid: 		testUser.Id,
		Name: 		testUser.Name,
		Surname:  	testUser.Surname,
		Email:      testUser.Email,
		Password:   testUser.Password,
	}

	protoNothing := &proto.Nothing{Dummy: true}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().AddUser(protoUser).Return(protoNothing, nil)

	output, err := ud.AddUser(ctx, protoUser)

	assert.Equal(t, protoNothing, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_CreateFolder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoFolder := &proto.Folder{
		Uid: testFolder.Uid,
		Name: testFolder.Name,
		Type: testFolder.Type,
	}

	protoNothing := &proto.Nothing{Dummy: true}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().CreateFolder(protoFolder).Return(protoNothing, nil)

	output, err := ud.CreateFolder(ctx, protoFolder)

	assert.Equal(t, protoNothing, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_GenerateSID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoNothing := &proto.Nothing{Dummy: true}

	protoSid := &proto.Sid{
		Sid: testSession.Id,
	}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().GenerateSID(protoNothing).Return(protoSid, nil)

	output, err := ud.GenerateSID(ctx, protoNothing)

	assert.Equal(t, protoSid, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_GenerateUID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoNothing := &proto.Nothing{Dummy: true}

	protoUid := &proto.Uid{
		Uid: testFolder.Uid,
	}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().GenerateUID(protoNothing).Return(protoUid, nil)

	output, err := ud.GenerateUID(ctx, protoNothing)

	assert.Equal(t, protoUid, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_GetFolderId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoFolder := &proto.Folder{
		Uid: testFolder.Uid,
		Name: testFolder.Name,
		Type: testFolder.Type,
	}

	protoFolderID := &proto.FolderId{
		Id: testFolder.Id,
	}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().GetFolderId(protoFolder).Return(protoFolderID, nil)

	output, err := ud.GetFolderId(ctx, protoFolder)

	assert.Equal(t, protoFolderID, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_GetFoldersList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoFolderUidType := &proto.FolderUidType{
		Uid: testFolder.Uid,
		Type: testFolder.Type,
	}

	protoFolderList := &proto.FolderList{
		Res: []*proto.FolderNameType{},
	}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().GetFoldersList(protoFolderUidType).Return(protoFolderList, nil)

	output, err := ud.GetFoldersList(ctx, protoFolderUidType)

	assert.Equal(t, protoFolderList, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_GetSessionByUID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoUid := &proto.Uid{
		Uid: testFolder.Uid,
	}

	protoSid := &proto.Sid{
		Sid: testSession.Id,
	}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().GetSessionByUID(protoUid).Return(protoSid, nil)

	output, err := ud.GetSessionByUID(ctx, protoUid)

	assert.Equal(t, protoSid, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoEmail := &proto.Email{
		Email: testUser.Email,
	}

	protoUser := &proto.User{
		Uid: 		testUser.Id,
		Name: 		testUser.Name,
		Surname:  	testUser.Surname,
		Email:      testUser.Email,
		Password:   testUser.Password,
	}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().GetUserByEmail(protoEmail).Return(protoUser, nil)

	output, err := ud.GetUserByEmail(ctx, protoEmail)

	assert.Equal(t, protoUser, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_GetUserByUID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoUid := &proto.Uid{
		Uid: testFolder.Uid,
	}

	protoUser := &proto.User{
		Uid: 		testUser.Id,
		Name: 		testUser.Name,
		Surname:  	testUser.Surname,
		Email:      testUser.Email,
		Password:   testUser.Password,
	}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().GetUserByUID(protoUid).Return(protoUser, nil)

	output, err := ud.GetUserByUID(ctx, protoUid)

	assert.Equal(t, protoUser, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_IsEmailExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoEmail := &proto.Email{
		Email: testUser.Email,
	}


	protoNothing := &proto.Nothing{Dummy: true}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().IsEmailExists(protoEmail).Return(protoNothing, nil)

	output, err := ud.IsEmailExists(ctx, protoEmail)

	assert.Equal(t, protoNothing, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_IsOkSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoSid := &proto.Sid{
		Sid: testSession.Id,
	}

	protoUid := &proto.Uid{
		Uid: testFolder.Uid,
	}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().IsOkSession(protoSid).Return(protoUid, nil)

	output, err := ud.IsOkSession(ctx, protoSid)

	assert.Equal(t, protoUid, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_RemoveFolder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoFolder := &proto.Folder{
		Uid: testFolder.Uid,
		Name: testFolder.Name,
		Type: testFolder.Type,
	}

	protoFolderID := &proto.FolderId{
		Id: testFolder.Id,
	}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().RemoveFolder(protoFolder).Return(protoFolderID, nil)

	output, err := ud.RemoveFolder(ctx, protoFolder)

	assert.Equal(t, protoFolderID, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_RemoveSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoSid := &proto.Sid{
		Sid: testSession.Id,
	}

	protoUid := &proto.Uid{
		Uid: testFolder.Uid,
	}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().RemoveSession(protoSid).Return(protoUid, nil)

	output, err := ud.RemoveSession(ctx, protoSid)

	assert.Equal(t, protoUid, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_RenameFolder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoRenameFolderMsg := &proto.RenameFolderMsg{
		Uid: testFolder.Id,
		Type: testFolder.Type,
		OldName: testFolder.Name,
		NewName: testFolder.Name,
	}

	protoNothing := &proto.Nothing{Dummy: true}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().RenameFolder(protoRenameFolderMsg).Return(protoNothing, nil)

	output, err := ud.RenameFolder(ctx, protoRenameFolderMsg)

	assert.Equal(t, protoNothing, output)
	assert.Equal(t, nil, err)
}

func TestUserManager_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	protoUser := &proto.User{
		Uid: 		testUser.Id,
		Name: 		testUser.Name,
		Surname:  	testUser.Surname,
		Email:      testUser.Email,
		Password:   testUser.Password,
	}

	protoUpdateProfileMsg := &proto.UpdateProfileMsg{
		NewUser: protoUser,
		Email: testUser.Email,
	}

	protoNothing := &proto.Nothing{Dummy: true}

	ctx := context.Background()

	mockUC := mock.NewMockInterface(ctrl)
	ud := New(mockUC)

	mockUC.EXPECT().UpdateProfile(protoUpdateProfileMsg).Return(protoNothing, nil)

	output, err := ud.UpdateProfile(ctx, protoUpdateProfileMsg)

	assert.Equal(t, protoNothing, output)
	assert.Equal(t, nil, err)
}
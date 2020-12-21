package UserUseCase
//
//import (
//	"Mailer/UserService/internal/UserModel"
//	mock "Mailer/UserService/internal/mocks"
//	proto "Mailer/UserService/proto"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//var (
//	testUser = UserModel.User{
//		Id:      	1,
//		Name: 		"UserName",
//		Surname:  	"UserSurname",
//		Email:      "UserEmail",
//		Password:   "UserPassword",
//		Img:     	"",
//	}
//
//	testSession = UserModel.Session{
//		Id:        	"SessionId",
//		UserId: 	1,
//		User:		&testUser,
//	}
//
//	testFolder = UserModel.Folder{
//		Uid:		1,
//		Type: 		"FolderType",
//		Name:		"FolderName",
//	}
//)
//
//func TestUseCase_AddSession(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoUser := &proto.User{
//		Uid: 		testUser.Id,
//		Name: 		testUser.Name,
//		Surname:  	testUser.Surname,
//		Email:      testUser.Email,
//		Password:   testUser.Password,
//	}
//
//	protoSessionMsg := &proto.AddSessionMsg {
//		Sid: testSession.Id,
//		User: protoUser,
//	}
//
//	protoNothing := &proto.Nothing{Dummy: true}
//
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().AddSession(testSession.Id, testUser.Id, &testUser).Return(nil)
//
//	output, err := uUC.AddSession(protoSessionMsg)
//
//	assert.Equal(t, protoNothing, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_AddUser(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoUser := &proto.User{
//		Uid: 		testUser.Id,
//		Name: 		testUser.Name,
//		Surname:  	testUser.Surname,
//		Email:      testUser.Email,
//		Password:   testUser.Password,
//	}
//
//	protoNothing := &proto.Nothing{Dummy: true}
//
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().AddUser(&testUser).Return(nil)
//
//	output, err := uUC.AddUser(protoUser)
//
//	assert.Equal(t, protoNothing, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_CreateFolder(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoFolder := &proto.Folder{
//		Type: 		testFolder.Type,
//		Uid: 		testFolder.Uid,
//		Name: 		testFolder.Name,
//	}
//
//	protoNothing := &proto.Nothing{Dummy: true}
//
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().CreateFolder(testFolder.Name, testFolder.Type, testFolder.Uid).Return(nil)
//
//	output, err := uUC.CreateFolder(protoFolder)
//
//	assert.Equal(t, protoNothing, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_GenerateSID(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoNothing := &proto.Nothing{Dummy: true}
//	protoSid := &proto.Sid{Sid: string("sid")}
//	sidDB := []rune("sid")
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().GenerateSID().Return(sidDB, nil)
//
//	output, err := uUC.GenerateSID(protoNothing)
//
//	assert.Equal(t, protoSid, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_GenerateUID(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoNothing := &proto.Nothing{Dummy: true}
//	protoUid := &proto.Uid{Uid: uint64(1)}
//	var uidDB = uint64(1)
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().GenerateUID().Return(uidDB, nil)
//
//	output, err := uUC.GenerateUID(protoNothing)
//
//	assert.Equal(t, protoUid, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_GetFolderId(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoFolder := &proto.Folder{
//		Type: 		testFolder.Type,
//		Uid: 		testFolder.Uid,
//		Name: 		testFolder.Name,
//	}
//
//	protoFolderId := &proto.FolderId{
//		Id: testFolder.Id,
//	}
//
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().GetFolderId(testFolder.Uid, testFolder.Type, testFolder.Name).Return(testFolder.Id, nil)
//
//	output, err := uUC.GetFolderId(protoFolder)
//
//	assert.Equal(t, protoFolderId, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_GetFoldersList(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoFolderUidType:= &proto.FolderUidType{
//		Uid: testFolder.Uid,
//		Type: testFolder.Type,
//	}
//
//	protoFolderList := &proto.FolderList{}
//	folderArr := []*UserModel.Folder{}
//
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().GetFoldersList(protoFolderUidType.Uid, protoFolderUidType.Type).Return(folderArr, nil)
//
//	output, err := uUC.GetFoldersList(protoFolderUidType)
//
//	assert.Equal(t, protoFolderList, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_GetSessionByUID(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoUid := &proto.Uid{Uid: uint64(1)}
//	protoSid := &proto.Sid{Sid: ""}
//
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().GetSessionByUID(protoUid.Uid).Return("", nil)
//
//	output, err := uUC.GetSessionByUID(protoUid)
//
//	assert.Equal(t, protoSid, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_GetUserByEmail(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoEmail := &proto.Email{
//		Email: testUser.Email,
//	}
//
//	protoUser := &proto.User{
//		Uid: 		testUser.Id,
//		Name: 		testUser.Name,
//		Surname:  	testUser.Surname,
//		Email:      testUser.Email,
//		Password:   testUser.Password,
//	}
//
//
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().GetUserByEmail(protoEmail.Email).Return(&testUser, nil)
//
//	output, err := uUC.GetUserByEmail(protoEmail)
//
//	assert.Equal(t, protoUser, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_GetUserByUID(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoUid := &proto.Uid{Uid: uint64(1)}
//
//	protoUser := &proto.User{
//		Uid: 		testUser.Id,
//		Name: 		testUser.Name,
//		Surname:  	testUser.Surname,
//		Email:      testUser.Email,
//		Password:   testUser.Password,
//	}
//
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().GetUserByUID(protoUid.Uid).Return(&testUser, nil)
//
//	output, err := uUC.GetUserByUID(protoUid)
//
//	assert.Equal(t, protoUser, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_IsEmailExists(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoEmail := &proto.Email{
//		Email: testUser.Email,
//	}
//
//	protoNothing := &proto.Nothing{Dummy: true}
//
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().IsEmailExists(protoEmail.Email).Return(nil)
//
//	output, err := uUC.IsEmailExists(protoEmail)
//
//	assert.Equal(t, protoNothing, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_IsOkSession(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoSid := &proto.Sid{Sid: string("sid")}
//	protoUid := &proto.Uid{Uid: uint64(1)}
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().IsOkSession(protoSid.Sid).Return(uint64(1), nil)
//
//	output, err := uUC.IsOkSession(protoSid)
//
//	assert.Equal(t, protoUid, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_RemoveFolder(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoFolder := &proto.Folder{
//		Type: 		testFolder.Type,
//		Uid: 		testFolder.Uid,
//		Name: 		testFolder.Name,
//	}
//
//	protoFolderId := &proto.FolderId{
//		Id: testFolder.Id,
//	}
//
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().GetFolderId(testFolder.Uid, testFolder.Type, testFolder.Name).Return(testFolder.Id, nil)
//	mockUserDB.EXPECT().RemoveFolder(testFolder.Id).Return(nil)
//
//	output, err := uUC.RemoveFolder(protoFolder)
//
//	assert.Equal(t, protoFolderId, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_RemoveSession(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoSid := &proto.Sid{Sid: string("sid")}
//	protoUid := &proto.Uid{Uid: uint64(1)}
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().RemoveSession(protoSid.Sid).Return(nil, uint64(1))
//
//	output, err := uUC.RemoveSession(protoSid)
//
//	assert.Equal(t, protoUid, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_RenameFolder(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoRenameFolderMsg := &proto.RenameFolderMsg{
//		Uid: testFolder.Uid,
//		Type: testFolder.Type,
//		OldName: testFolder.Name,
//		NewName: testFolder.Name,
//	}
//	protoNothing := &proto.Nothing{Dummy: true}
//
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().RenameFolder(
//		protoRenameFolderMsg.Uid,
//		protoRenameFolderMsg.Type,
//		protoRenameFolderMsg.OldName,
//		protoRenameFolderMsg.NewName).Return(nil)
//
//	output, err := uUC.RenameFolder(protoRenameFolderMsg)
//
//	assert.Equal(t, protoNothing, output)
//	assert.Equal(t, nil, err)
//}
//
//func TestUseCase_UpdateProfile(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	protoUser := &proto.User{
//		Uid: 		testUser.Id,
//		Name: 		testUser.Name,
//		Surname:  	testUser.Surname,
//		Email:      testUser.Email,
//		Password:   testUser.Password,
//	}
//
//	protoUpdateProfileMsg := &proto.UpdateProfileMsg {
//		Email: testUser.Email,
//		NewUser: protoUser,
//	}
//
//	protoNothing := &proto.Nothing{Dummy: true}
//
//	mockUserDB := mock.NewMockUserDB(ctrl)
//
//	uUC := New(mockUserDB)
//
//	mockUserDB.EXPECT().UpdateProfile(testUser, protoUpdateProfileMsg.Email).Return(nil)
//
//	output, err := uUC.UpdateProfile(protoUpdateProfileMsg)
//
//	assert.Equal(t, protoNothing, output)
//	assert.Equal(t, nil, err)
//}
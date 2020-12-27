// Code generated by MockGen. DO NOT EDIT.
// Source: ./DataBaseRequests.go

// Package mocks is a generated GoMock package.
package mocks

import (
	UserModel "Mailer/UserService/internal/UserModel"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserDB is a mock of UserDB interface
type MockUserDB struct {
	ctrl     *gomock.Controller
	recorder *MockUserDBMockRecorder
}

// MockUserDBMockRecorder is the mock recorder for MockUserDB
type MockUserDBMockRecorder struct {
	mock *MockUserDB
}

// NewMockUserDB creates a new mock instance
func NewMockUserDB(ctrl *gomock.Controller) *MockUserDB {
	mock := &MockUserDB{ctrl: ctrl}
	mock.recorder = &MockUserDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserDB) EXPECT() *MockUserDBMockRecorder {
	return m.recorder
}

// IsEmailExists mocks base method
func (m *MockUserDB) IsEmailExists(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmailExists", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsEmailExists indicates an expected call of IsEmailExists
func (mr *MockUserDBMockRecorder) IsEmailExists(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmailExists", reflect.TypeOf((*MockUserDB)(nil).IsEmailExists), email)
}

// AddSession mocks base method
func (m *MockUserDB) AddSession(sid string, uid uint64, user *UserModel.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSession", sid, uid, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSession indicates an expected call of AddSession
func (mr *MockUserDBMockRecorder) AddSession(sid, uid, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSession", reflect.TypeOf((*MockUserDB)(nil).AddSession), sid, uid, user)
}

// AddUser mocks base method
func (m *MockUserDB) AddUser(user *UserModel.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser
func (mr *MockUserDBMockRecorder) AddUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserDB)(nil).AddUser), user)
}

// GenerateSID mocks base method
func (m *MockUserDB) GenerateSID() ([]rune, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSID")
	ret0, _ := ret[0].([]rune)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateSID indicates an expected call of GenerateSID
func (mr *MockUserDBMockRecorder) GenerateSID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSID", reflect.TypeOf((*MockUserDB)(nil).GenerateSID))
}

// GenerateUID mocks base method
func (m *MockUserDB) GenerateUID() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUID")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateUID indicates an expected call of GenerateUID
func (mr *MockUserDBMockRecorder) GenerateUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUID", reflect.TypeOf((*MockUserDB)(nil).GenerateUID))
}

// GetUserByEmail mocks base method
func (m *MockUserDB) GetUserByEmail(email string) (*UserModel.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(*UserModel.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail
func (mr *MockUserDBMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserDB)(nil).GetUserByEmail), email)
}

// GetUserByUID mocks base method
func (m *MockUserDB) GetUserByUID(uid uint64) (*UserModel.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUID", uid)
	ret0, _ := ret[0].(*UserModel.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUID indicates an expected call of GetUserByUID
func (mr *MockUserDBMockRecorder) GetUserByUID(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUID", reflect.TypeOf((*MockUserDB)(nil).GetUserByUID), uid)
}

// IsOkSession mocks base method
func (m *MockUserDB) IsOkSession(sid string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOkSession", sid)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOkSession indicates an expected call of IsOkSession
func (mr *MockUserDBMockRecorder) IsOkSession(sid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOkSession", reflect.TypeOf((*MockUserDB)(nil).IsOkSession), sid)
}

// UpdateProfile mocks base method
func (m *MockUserDB) UpdateProfile(newUser UserModel.User, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", newUser, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile
func (mr *MockUserDBMockRecorder) UpdateProfile(newUser, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserDB)(nil).UpdateProfile), newUser, email)
}

// RemoveSession mocks base method
func (m *MockUserDB) RemoveSession(sid string) (error, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSession", sid)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// RemoveSession indicates an expected call of RemoveSession
func (mr *MockUserDBMockRecorder) RemoveSession(sid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSession", reflect.TypeOf((*MockUserDB)(nil).RemoveSession), sid)
}

// GetSessionByUID mocks base method
func (m *MockUserDB) GetSessionByUID(uid uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionByUID", uid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionByUID indicates an expected call of GetSessionByUID
func (mr *MockUserDBMockRecorder) GetSessionByUID(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByUID", reflect.TypeOf((*MockUserDB)(nil).GetSessionByUID), uid)
}

// GetFolderId mocks base method
func (m *MockUserDB) GetFolderId(uid uint64, kind, name string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFolderId", uid, kind, name)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFolderId indicates an expected call of GetFolderId
func (mr *MockUserDBMockRecorder) GetFolderId(uid, kind, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFolderId", reflect.TypeOf((*MockUserDB)(nil).GetFolderId), uid, kind, name)
}

// CreateFolder mocks base method
func (m *MockUserDB) CreateFolder(name, kind string, uid uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFolder", name, kind, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFolder indicates an expected call of CreateFolder
func (mr *MockUserDBMockRecorder) CreateFolder(name, kind, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFolder", reflect.TypeOf((*MockUserDB)(nil).CreateFolder), name, kind, uid)
}

// RenameFolder mocks base method
func (m *MockUserDB) RenameFolder(uid uint64, kind, oldName, newName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameFolder", uid, kind, oldName, newName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameFolder indicates an expected call of RenameFolder
func (mr *MockUserDBMockRecorder) RenameFolder(uid, kind, oldName, newName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameFolder", reflect.TypeOf((*MockUserDB)(nil).RenameFolder), uid, kind, oldName, newName)
}

// RemoveFolder mocks base method
func (m *MockUserDB) RemoveFolder(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFolder", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFolder indicates an expected call of RemoveFolder
func (mr *MockUserDBMockRecorder) RemoveFolder(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFolder", reflect.TypeOf((*MockUserDB)(nil).RemoveFolder), id)
}

// GetFoldersList mocks base method
func (m *MockUserDB) GetFoldersList(uid uint64, kind string) ([]*UserModel.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFoldersList", uid, kind)
	ret0, _ := ret[0].([]*UserModel.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFoldersList indicates an expected call of GetFoldersList
func (mr *MockUserDBMockRecorder) GetFoldersList(uid, kind interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFoldersList", reflect.TypeOf((*MockUserDB)(nil).GetFoldersList), uid, kind)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: ./mail.pb.go

// Package mock_letterService is a generated GoMock package.
package mock_letterService

import (
	letterService "Mailer/MailService/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockLetterServiceClient is a mock of LetterServiceClient interface
type MockLetterServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockLetterServiceClientMockRecorder
}

// MockLetterServiceClientMockRecorder is the mock recorder for MockLetterServiceClient
type MockLetterServiceClientMockRecorder struct {
	mock *MockLetterServiceClient
}

// NewMockLetterServiceClient creates a new mock instance
func NewMockLetterServiceClient(ctrl *gomock.Controller) *MockLetterServiceClient {
	mock := &MockLetterServiceClient{ctrl: ctrl}
	mock.recorder = &MockLetterServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLetterServiceClient) EXPECT() *MockLetterServiceClientMockRecorder {
	return m.recorder
}

// GetLettersByDirRecv mocks base method
func (m *MockLetterServiceClient) GetLettersByDirRecv(ctx context.Context, in *letterService.DirName, opts ...grpc.CallOption) (*letterService.LetterListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLettersByDirRecv", varargs...)
	ret0, _ := ret[0].(*letterService.LetterListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLettersByDirRecv indicates an expected call of GetLettersByDirRecv
func (mr *MockLetterServiceClientMockRecorder) GetLettersByDirRecv(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLettersByDirRecv", reflect.TypeOf((*MockLetterServiceClient)(nil).GetLettersByDirRecv), varargs...)
}

// GetLettersByDirSend mocks base method
func (m *MockLetterServiceClient) GetLettersByDirSend(ctx context.Context, in *letterService.DirName, opts ...grpc.CallOption) (*letterService.LetterListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLettersByDirSend", varargs...)
	ret0, _ := ret[0].(*letterService.LetterListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLettersByDirSend indicates an expected call of GetLettersByDirSend
func (mr *MockLetterServiceClientMockRecorder) GetLettersByDirSend(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLettersByDirSend", reflect.TypeOf((*MockLetterServiceClient)(nil).GetLettersByDirSend), varargs...)
}

// GetLettersRecv mocks base method
func (m *MockLetterServiceClient) GetLettersRecv(ctx context.Context, in *letterService.Email, opts ...grpc.CallOption) (*letterService.LetterListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLettersRecv", varargs...)
	ret0, _ := ret[0].(*letterService.LetterListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLettersRecv indicates an expected call of GetLettersRecv
func (mr *MockLetterServiceClientMockRecorder) GetLettersRecv(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLettersRecv", reflect.TypeOf((*MockLetterServiceClient)(nil).GetLettersRecv), varargs...)
}

// GetLettersSend mocks base method
func (m *MockLetterServiceClient) GetLettersSend(ctx context.Context, in *letterService.Email, opts ...grpc.CallOption) (*letterService.LetterListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLettersSend", varargs...)
	ret0, _ := ret[0].(*letterService.LetterListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLettersSend indicates an expected call of GetLettersSend
func (mr *MockLetterServiceClientMockRecorder) GetLettersSend(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLettersSend", reflect.TypeOf((*MockLetterServiceClient)(nil).GetLettersSend), varargs...)
}

// SaveLetter mocks base method
func (m *MockLetterServiceClient) SaveLetter(ctx context.Context, in *letterService.Letter, opts ...grpc.CallOption) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SaveLetter", varargs...)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveLetter indicates an expected call of SaveLetter
func (mr *MockLetterServiceClientMockRecorder) SaveLetter(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLetter", reflect.TypeOf((*MockLetterServiceClient)(nil).SaveLetter), varargs...)
}

// WatchedLetter mocks base method
func (m *MockLetterServiceClient) WatchedLetter(ctx context.Context, in *letterService.Lid, opts ...grpc.CallOption) (*letterService.LetterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WatchedLetter", varargs...)
	ret0, _ := ret[0].(*letterService.LetterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchedLetter indicates an expected call of WatchedLetter
func (mr *MockLetterServiceClientMockRecorder) WatchedLetter(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchedLetter", reflect.TypeOf((*MockLetterServiceClient)(nil).WatchedLetter), varargs...)
}

// RemoveLetter mocks base method
func (m *MockLetterServiceClient) RemoveLetter(ctx context.Context, in *letterService.Lid, opts ...grpc.CallOption) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveLetter", varargs...)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveLetter indicates an expected call of RemoveLetter
func (mr *MockLetterServiceClientMockRecorder) RemoveLetter(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLetter", reflect.TypeOf((*MockLetterServiceClient)(nil).RemoveLetter), varargs...)
}

// AddLetterToDir mocks base method
func (m *MockLetterServiceClient) AddLetterToDir(ctx context.Context, in *letterService.DirLid, opts ...grpc.CallOption) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddLetterToDir", varargs...)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLetterToDir indicates an expected call of AddLetterToDir
func (mr *MockLetterServiceClientMockRecorder) AddLetterToDir(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLetterToDir", reflect.TypeOf((*MockLetterServiceClient)(nil).AddLetterToDir), varargs...)
}

// RemoveLetterFromDir mocks base method
func (m *MockLetterServiceClient) RemoveLetterFromDir(ctx context.Context, in *letterService.DirLid, opts ...grpc.CallOption) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveLetterFromDir", varargs...)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveLetterFromDir indicates an expected call of RemoveLetterFromDir
func (mr *MockLetterServiceClientMockRecorder) RemoveLetterFromDir(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLetterFromDir", reflect.TypeOf((*MockLetterServiceClient)(nil).RemoveLetterFromDir), varargs...)
}

// RemoveDir mocks base method
func (m *MockLetterServiceClient) RemoveDir(ctx context.Context, in *letterService.DirLid, opts ...grpc.CallOption) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveDir", varargs...)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveDir indicates an expected call of RemoveDir
func (mr *MockLetterServiceClientMockRecorder) RemoveDir(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDir", reflect.TypeOf((*MockLetterServiceClient)(nil).RemoveDir), varargs...)
}

// FindSimilar mocks base method
func (m *MockLetterServiceClient) FindSimilar(ctx context.Context, in *letterService.Similar, opts ...grpc.CallOption) (*letterService.SimRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindSimilar", varargs...)
	ret0, _ := ret[0].(*letterService.SimRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSimilar indicates an expected call of FindSimilar
func (mr *MockLetterServiceClientMockRecorder) FindSimilar(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSimilar", reflect.TypeOf((*MockLetterServiceClient)(nil).FindSimilar), varargs...)
}

// GetLetterBy mocks base method
func (m *MockLetterServiceClient) GetLetterBy(ctx context.Context, in *letterService.GetBy, opts ...grpc.CallOption) (*letterService.LetterListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLetterBy", varargs...)
	ret0, _ := ret[0].(*letterService.LetterListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLetterBy indicates an expected call of GetLetterBy
func (mr *MockLetterServiceClientMockRecorder) GetLetterBy(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLetterBy", reflect.TypeOf((*MockLetterServiceClient)(nil).GetLetterBy), varargs...)
}

// SetLetterInSpam mocks base method
func (m *MockLetterServiceClient) SetLetterInSpam(ctx context.Context, in *letterService.Lid, opts ...grpc.CallOption) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetLetterInSpam", varargs...)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetLetterInSpam indicates an expected call of SetLetterInSpam
func (mr *MockLetterServiceClientMockRecorder) SetLetterInSpam(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLetterInSpam", reflect.TypeOf((*MockLetterServiceClient)(nil).SetLetterInSpam), varargs...)
}

// SetLetterInBox mocks base method
func (m *MockLetterServiceClient) SetLetterInBox(ctx context.Context, in *letterService.Lid, opts ...grpc.CallOption) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetLetterInBox", varargs...)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetLetterInBox indicates an expected call of SetLetterInBox
func (mr *MockLetterServiceClientMockRecorder) SetLetterInBox(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLetterInBox", reflect.TypeOf((*MockLetterServiceClient)(nil).SetLetterInBox), varargs...)
}

// MockLetterServiceServer is a mock of LetterServiceServer interface
type MockLetterServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockLetterServiceServerMockRecorder
}

// MockLetterServiceServerMockRecorder is the mock recorder for MockLetterServiceServer
type MockLetterServiceServerMockRecorder struct {
	mock *MockLetterServiceServer
}

// NewMockLetterServiceServer creates a new mock instance
func NewMockLetterServiceServer(ctrl *gomock.Controller) *MockLetterServiceServer {
	mock := &MockLetterServiceServer{ctrl: ctrl}
	mock.recorder = &MockLetterServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLetterServiceServer) EXPECT() *MockLetterServiceServerMockRecorder {
	return m.recorder
}

// GetLettersByDirRecv mocks base method
func (m *MockLetterServiceServer) GetLettersByDirRecv(arg0 context.Context, arg1 *letterService.DirName) (*letterService.LetterListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLettersByDirRecv", arg0, arg1)
	ret0, _ := ret[0].(*letterService.LetterListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLettersByDirRecv indicates an expected call of GetLettersByDirRecv
func (mr *MockLetterServiceServerMockRecorder) GetLettersByDirRecv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLettersByDirRecv", reflect.TypeOf((*MockLetterServiceServer)(nil).GetLettersByDirRecv), arg0, arg1)
}

// GetLettersByDirSend mocks base method
func (m *MockLetterServiceServer) GetLettersByDirSend(arg0 context.Context, arg1 *letterService.DirName) (*letterService.LetterListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLettersByDirSend", arg0, arg1)
	ret0, _ := ret[0].(*letterService.LetterListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLettersByDirSend indicates an expected call of GetLettersByDirSend
func (mr *MockLetterServiceServerMockRecorder) GetLettersByDirSend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLettersByDirSend", reflect.TypeOf((*MockLetterServiceServer)(nil).GetLettersByDirSend), arg0, arg1)
}

// GetLettersRecv mocks base method
func (m *MockLetterServiceServer) GetLettersRecv(arg0 context.Context, arg1 *letterService.Email) (*letterService.LetterListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLettersRecv", arg0, arg1)
	ret0, _ := ret[0].(*letterService.LetterListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLettersRecv indicates an expected call of GetLettersRecv
func (mr *MockLetterServiceServerMockRecorder) GetLettersRecv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLettersRecv", reflect.TypeOf((*MockLetterServiceServer)(nil).GetLettersRecv), arg0, arg1)
}

// GetLettersSend mocks base method
func (m *MockLetterServiceServer) GetLettersSend(arg0 context.Context, arg1 *letterService.Email) (*letterService.LetterListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLettersSend", arg0, arg1)
	ret0, _ := ret[0].(*letterService.LetterListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLettersSend indicates an expected call of GetLettersSend
func (mr *MockLetterServiceServerMockRecorder) GetLettersSend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLettersSend", reflect.TypeOf((*MockLetterServiceServer)(nil).GetLettersSend), arg0, arg1)
}

// SaveLetter mocks base method
func (m *MockLetterServiceServer) SaveLetter(arg0 context.Context, arg1 *letterService.Letter) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLetter", arg0, arg1)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveLetter indicates an expected call of SaveLetter
func (mr *MockLetterServiceServerMockRecorder) SaveLetter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLetter", reflect.TypeOf((*MockLetterServiceServer)(nil).SaveLetter), arg0, arg1)
}

// WatchedLetter mocks base method
func (m *MockLetterServiceServer) WatchedLetter(arg0 context.Context, arg1 *letterService.Lid) (*letterService.LetterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchedLetter", arg0, arg1)
	ret0, _ := ret[0].(*letterService.LetterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchedLetter indicates an expected call of WatchedLetter
func (mr *MockLetterServiceServerMockRecorder) WatchedLetter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchedLetter", reflect.TypeOf((*MockLetterServiceServer)(nil).WatchedLetter), arg0, arg1)
}

// RemoveLetter mocks base method
func (m *MockLetterServiceServer) RemoveLetter(arg0 context.Context, arg1 *letterService.Lid) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLetter", arg0, arg1)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveLetter indicates an expected call of RemoveLetter
func (mr *MockLetterServiceServerMockRecorder) RemoveLetter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLetter", reflect.TypeOf((*MockLetterServiceServer)(nil).RemoveLetter), arg0, arg1)
}

// AddLetterToDir mocks base method
func (m *MockLetterServiceServer) AddLetterToDir(arg0 context.Context, arg1 *letterService.DirLid) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLetterToDir", arg0, arg1)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLetterToDir indicates an expected call of AddLetterToDir
func (mr *MockLetterServiceServerMockRecorder) AddLetterToDir(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLetterToDir", reflect.TypeOf((*MockLetterServiceServer)(nil).AddLetterToDir), arg0, arg1)
}

// RemoveLetterFromDir mocks base method
func (m *MockLetterServiceServer) RemoveLetterFromDir(arg0 context.Context, arg1 *letterService.DirLid) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLetterFromDir", arg0, arg1)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveLetterFromDir indicates an expected call of RemoveLetterFromDir
func (mr *MockLetterServiceServerMockRecorder) RemoveLetterFromDir(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLetterFromDir", reflect.TypeOf((*MockLetterServiceServer)(nil).RemoveLetterFromDir), arg0, arg1)
}

// RemoveDir mocks base method
func (m *MockLetterServiceServer) RemoveDir(arg0 context.Context, arg1 *letterService.DirLid) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDir", arg0, arg1)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveDir indicates an expected call of RemoveDir
func (mr *MockLetterServiceServerMockRecorder) RemoveDir(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDir", reflect.TypeOf((*MockLetterServiceServer)(nil).RemoveDir), arg0, arg1)
}

// FindSimilar mocks base method
func (m *MockLetterServiceServer) FindSimilar(arg0 context.Context, arg1 *letterService.Similar) (*letterService.SimRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSimilar", arg0, arg1)
	ret0, _ := ret[0].(*letterService.SimRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSimilar indicates an expected call of FindSimilar
func (mr *MockLetterServiceServerMockRecorder) FindSimilar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSimilar", reflect.TypeOf((*MockLetterServiceServer)(nil).FindSimilar), arg0, arg1)
}

// GetLetterBy mocks base method
func (m *MockLetterServiceServer) GetLetterBy(arg0 context.Context, arg1 *letterService.GetBy) (*letterService.LetterListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLetterBy", arg0, arg1)
	ret0, _ := ret[0].(*letterService.LetterListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLetterBy indicates an expected call of GetLetterBy
func (mr *MockLetterServiceServerMockRecorder) GetLetterBy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLetterBy", reflect.TypeOf((*MockLetterServiceServer)(nil).GetLetterBy), arg0, arg1)
}

// SetLetterInSpam mocks base method
func (m *MockLetterServiceServer) SetLetterInSpam(arg0 context.Context, arg1 *letterService.Lid) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLetterInSpam", arg0, arg1)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetLetterInSpam indicates an expected call of SetLetterInSpam
func (mr *MockLetterServiceServerMockRecorder) SetLetterInSpam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLetterInSpam", reflect.TypeOf((*MockLetterServiceServer)(nil).SetLetterInSpam), arg0, arg1)
}

// SetLetterInBox mocks base method
func (m *MockLetterServiceServer) SetLetterInBox(arg0 context.Context, arg1 *letterService.Lid) (*letterService.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLetterInBox", arg0, arg1)
	ret0, _ := ret[0].(*letterService.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetLetterInBox indicates an expected call of SetLetterInBox
func (mr *MockLetterServiceServerMockRecorder) SetLetterInBox(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLetterInBox", reflect.TypeOf((*MockLetterServiceServer)(nil).SetLetterInBox), arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: Letter.go

// Package mock_LetterUseCase is a generated GoMock package.
package mock_LetterUseCase

import (
	LetterModel "Mailer/MainApplication/internal/Letter/LetterModel"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLetterUseCase is a mock of LetterUseCase interface
type MockLetterUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockLetterUseCaseMockRecorder
}

// MockLetterUseCaseMockRecorder is the mock recorder for MockLetterUseCase
type MockLetterUseCaseMockRecorder struct {
	mock *MockLetterUseCase
}

// NewMockLetterUseCase creates a new mock instance
func NewMockLetterUseCase(ctrl *gomock.Controller) *MockLetterUseCase {
	mock := &MockLetterUseCase{ctrl: ctrl}
	mock.recorder = &MockLetterUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLetterUseCase) EXPECT() *MockLetterUseCaseMockRecorder {
	return m.recorder
}

// SaveLetter mocks base method
func (m *MockLetterUseCase) SaveLetter(letter *LetterModel.Letter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLetter", letter)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLetter indicates an expected call of SaveLetter
func (mr *MockLetterUseCaseMockRecorder) SaveLetter(letter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLetter", reflect.TypeOf((*MockLetterUseCase)(nil).SaveLetter), letter)
}

// GetReceivedLetters mocks base method
func (m *MockLetterUseCase) GetReceivedLetters(email string) (error, []LetterModel.Letter) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceivedLetters", email)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].([]LetterModel.Letter)
	return ret0, ret1
}

// GetReceivedLetters indicates an expected call of GetReceivedLetters
func (mr *MockLetterUseCaseMockRecorder) GetReceivedLetters(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceivedLetters", reflect.TypeOf((*MockLetterUseCase)(nil).GetReceivedLetters), email)
}

// GetSendedLetters mocks base method
func (m *MockLetterUseCase) GetSendedLetters(email string) (error, []LetterModel.Letter) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSendedLetters", email)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].([]LetterModel.Letter)
	return ret0, ret1
}

// GetSendedLetters indicates an expected call of GetSendedLetters
func (mr *MockLetterUseCaseMockRecorder) GetSendedLetters(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSendedLetters", reflect.TypeOf((*MockLetterUseCase)(nil).GetSendedLetters), email)
}

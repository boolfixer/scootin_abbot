// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/user_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "main/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// GetByApiKey mocks base method.
func (m *MockUserRepository) GetByApiKey(apiKey string) (model.User, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByApiKey", apiKey)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetByApiKey indicates an expected call of GetByApiKey.
func (mr *MockUserRepositoryMockRecorder) GetByApiKey(apiKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByApiKey", reflect.TypeOf((*MockUserRepository)(nil).GetByApiKey), apiKey)
}

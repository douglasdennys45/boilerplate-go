// Code generated by MockGen. DO NOT EDIT.
// Source: lib/ports/repositories/add-log-repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	repositories "clean-architeture/lib/ports/repositories"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAddLogRepository is a mock of AddLogRepository interface
type MockAddLogRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAddLogRepositoryMockRecorder
}

// MockAddLogRepositoryMockRecorder is the mock recorder for MockAddLogRepository
type MockAddLogRepositoryMockRecorder struct {
	mock *MockAddLogRepository
}

// NewMockAddLogRepository creates a new mock instance
func NewMockAddLogRepository(ctrl *gomock.Controller) *MockAddLogRepository {
	mock := &MockAddLogRepository{ctrl: ctrl}
	mock.recorder = &MockAddLogRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAddLogRepository) EXPECT() *MockAddLogRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockAddLogRepository) Add(data repositories.AddLogRequest) repositories.AddLogResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", data)
	ret0, _ := ret[0].(repositories.AddLogResponse)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockAddLogRepositoryMockRecorder) Add(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockAddLogRepository)(nil).Add), data)
}
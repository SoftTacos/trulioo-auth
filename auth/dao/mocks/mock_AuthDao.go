// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/softtacos/trulioo-auth/auth/dao (interfaces: AuthDao)

// Package controller is a generated GoMock package.
package controller

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthDao is a mock of AuthDao interface.
type MockAuthDao struct {
	ctrl     *gomock.Controller
	recorder *MockAuthDaoMockRecorder
}

// MockAuthDaoMockRecorder is the mock recorder for MockAuthDao.
type MockAuthDaoMockRecorder struct {
	mock *MockAuthDao
}

// NewMockAuthDao creates a new mock instance.
func NewMockAuthDao(ctrl *gomock.Controller) *MockAuthDao {
	mock := &MockAuthDao{ctrl: ctrl}
	mock.recorder = &MockAuthDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthDao) EXPECT() *MockAuthDaoMockRecorder {
	return m.recorder
}

// CreatePassword mocks base method.
func (m *MockAuthDao) CreatePassword(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePassword indicates an expected call of CreatePassword.
func (mr *MockAuthDaoMockRecorder) CreatePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePassword", reflect.TypeOf((*MockAuthDao)(nil).CreatePassword), arg0, arg1)
}

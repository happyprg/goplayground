// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/happyprg/Documents/workspace_etc/goplayground/pattern/gomock/util/file_util.go

// Package util is a generated GoMock package.
package util

import (
	gomock "github.com/golang/mock/gomock"
	os "os"
	reflect "reflect"
)

// MockOsInterface is a mock of OsInterface interface
type MockOsInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOsInterfaceMockRecorder
}

// MockOsInterfaceMockRecorder is the mock recorder for MockOsInterface
type MockOsInterfaceMockRecorder struct {
	mock *MockOsInterface
}

// NewMockOsInterface creates a new mock instance
func NewMockOsInterface(ctrl *gomock.Controller) *MockOsInterface {
	mock := &MockOsInterface{ctrl: ctrl}
	mock.recorder = &MockOsInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOsInterface) EXPECT() *MockOsInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockOsInterface) Create(name string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOsInterfaceMockRecorder) Create(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOsInterface)(nil).Create), name)
}

// MockFileInterface is a mock of FileInterface interface
type MockFileInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFileInterfaceMockRecorder
}

// MockFileInterfaceMockRecorder is the mock recorder for MockFileInterface
type MockFileInterfaceMockRecorder struct {
	mock *MockFileInterface
}

// NewMockFileInterface creates a new mock instance
func NewMockFileInterface(ctrl *gomock.Controller) *MockFileInterface {
	mock := &MockFileInterface{ctrl: ctrl}
	mock.recorder = &MockFileInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileInterface) EXPECT() *MockFileInterfaceMockRecorder {
	return m.recorder
}

// WriteString mocks base method
func (m *MockFileInterface) WriteString(f *os.File, s string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteString", f, s)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteString indicates an expected call of WriteString
func (mr *MockFileInterfaceMockRecorder) WriteString(f, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteString", reflect.TypeOf((*MockFileInterface)(nil).WriteString), f, s)
}

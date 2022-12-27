// Code generated by MockGen. DO NOT EDIT.
// Source: ./storage/storage.go

// Package storage is a generated GoMock package.
package storage

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKeyValueStore is a mock of KeyValueStore interface.
type MockKeyValueStore struct {
	ctrl     *gomock.Controller
	recorder *MockKeyValueStoreMockRecorder
}

// MockKeyValueStoreMockRecorder is the mock recorder for MockKeyValueStore.
type MockKeyValueStoreMockRecorder struct {
	mock *MockKeyValueStore
}

// NewMockKeyValueStore creates a new mock instance.
func NewMockKeyValueStore(ctrl *gomock.Controller) *MockKeyValueStore {
	mock := &MockKeyValueStore{ctrl: ctrl}
	mock.recorder = &MockKeyValueStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyValueStore) EXPECT() *MockKeyValueStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockKeyValueStore) Get(arg0 Key) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKeyValueStoreMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKeyValueStore)(nil).Get), arg0)
}

// Set mocks base method.
func (m *MockKeyValueStore) Set(arg0 Key, arg1 Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockKeyValueStoreMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockKeyValueStore)(nil).Set), arg0, arg1)
}
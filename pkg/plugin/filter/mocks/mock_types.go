// Code generated by MockGen. DO NOT EDIT.
// Source: types.go
//
// Generated by this command:
//
//	mockgen -source=types.go -destination=mocks/mock_types.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	net "net"
	reflect "reflect"

	ebpf "github.com/cilium/ebpf"
	gomock "go.uber.org/mock/gomock"
)

// MockIFilterMap is a mock of IFilterMap interface.
type MockIFilterMap struct {
	ctrl     *gomock.Controller
	recorder *MockIFilterMapMockRecorder
}

// MockIFilterMapMockRecorder is the mock recorder for MockIFilterMap.
type MockIFilterMapMockRecorder struct {
	mock *MockIFilterMap
}

// NewMockIFilterMap creates a new mock instance.
func NewMockIFilterMap(ctrl *gomock.Controller) *MockIFilterMap {
	mock := &MockIFilterMap{ctrl: ctrl}
	mock.recorder = &MockIFilterMapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFilterMap) EXPECT() *MockIFilterMapMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIFilterMap) Add(arg0 []net.IP) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockIFilterMapMockRecorder) Add(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIFilterMap)(nil).Add), arg0)
}

// Close mocks base method.
func (m *MockIFilterMap) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockIFilterMapMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIFilterMap)(nil).Close))
}

// Delete mocks base method.
func (m *MockIFilterMap) Delete(arg0 []net.IP) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIFilterMapMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIFilterMap)(nil).Delete), arg0)
}

// MockIEbpfMap is a mock of IEbpfMap interface.
type MockIEbpfMap struct {
	ctrl     *gomock.Controller
	recorder *MockIEbpfMapMockRecorder
}

// MockIEbpfMapMockRecorder is the mock recorder for MockIEbpfMap.
type MockIEbpfMapMockRecorder struct {
	mock *MockIEbpfMap
}

// NewMockIEbpfMap creates a new mock instance.
func NewMockIEbpfMap(ctrl *gomock.Controller) *MockIEbpfMap {
	mock := &MockIEbpfMap{ctrl: ctrl}
	mock.recorder = &MockIEbpfMapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEbpfMap) EXPECT() *MockIEbpfMapMockRecorder {
	return m.recorder
}

// BatchDelete mocks base method.
func (m *MockIEbpfMap) BatchDelete(keys any, opts *ebpf.BatchOptions) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchDelete", keys, opts)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchDelete indicates an expected call of BatchDelete.
func (mr *MockIEbpfMapMockRecorder) BatchDelete(keys, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDelete", reflect.TypeOf((*MockIEbpfMap)(nil).BatchDelete), keys, opts)
}

// BatchUpdate mocks base method.
func (m *MockIEbpfMap) BatchUpdate(keys, values any, opts *ebpf.BatchOptions) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdate", keys, values, opts)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchUpdate indicates an expected call of BatchUpdate.
func (mr *MockIEbpfMapMockRecorder) BatchUpdate(keys, values, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdate", reflect.TypeOf((*MockIEbpfMap)(nil).BatchUpdate), keys, values, opts)
}

// Close mocks base method.
func (m *MockIEbpfMap) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIEbpfMapMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIEbpfMap)(nil).Close))
}

// Delete mocks base method.
func (m *MockIEbpfMap) Delete(key any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIEbpfMapMockRecorder) Delete(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIEbpfMap)(nil).Delete), key)
}

// Put mocks base method.
func (m *MockIEbpfMap) Put(key, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockIEbpfMapMockRecorder) Put(key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockIEbpfMap)(nil).Put), key, value)
}

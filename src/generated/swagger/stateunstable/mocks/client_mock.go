// Code generated by MockGen. DO NOT EDIT.
// Source: ./generated/swagger/stateunstable/client/operations/operations_client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	operations "github.com/chronosphereio/chronoctl-core/src/generated/swagger/stateunstable/client/operations"
	runtime "github.com/go-openapi/runtime"
	gomock "github.com/golang/mock/gomock"
)

// MockClientService is a mock of ClientService interface.
type MockClientService struct {
	ctrl     *gomock.Controller
	recorder *MockClientServiceMockRecorder
}

// MockClientServiceMockRecorder is the mock recorder for MockClientService.
type MockClientServiceMockRecorder struct {
	mock *MockClientService
}

// NewMockClientService creates a new mock instance.
func NewMockClientService(ctrl *gomock.Controller) *MockClientService {
	mock := &MockClientService{ctrl: ctrl}
	mock.recorder = &MockClientServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientService) EXPECT() *MockClientServiceMockRecorder {
	return m.recorder
}

// Echo mocks base method.
func (m *MockClientService) Echo(params *operations.EchoParams, opts ...operations.ClientOption) (*operations.EchoOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Echo", varargs...)
	ret0, _ := ret[0].(*operations.EchoOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Echo indicates an expected call of Echo.
func (mr *MockClientServiceMockRecorder) Echo(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Echo", reflect.TypeOf((*MockClientService)(nil).Echo), varargs...)
}

// SetTransport mocks base method.
func (m *MockClientService) SetTransport(transport runtime.ClientTransport) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTransport", transport)
}

// SetTransport indicates an expected call of SetTransport.
func (mr *MockClientServiceMockRecorder) SetTransport(transport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransport", reflect.TypeOf((*MockClientService)(nil).SetTransport), transport)
}

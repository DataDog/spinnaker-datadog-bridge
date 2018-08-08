// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bobbytables/spinnaker-datadog-bridge/spinnaker (interfaces: Handler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "github.com/DataDog/spinnaker-datadog-bridge/spinnaker/types"
	gomock "github.com/golang/mock/gomock"
)

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method
func (m *MockHandler) Handle(arg0 *types.IncomingWebhook) error {
	ret := m.ctrl.Call(m, "Handle", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle
func (mr *MockHandlerMockRecorder) Handle(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockHandler)(nil).Handle), arg0)
}

// Name mocks base method
func (m *MockHandler) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockHandlerMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockHandler)(nil).Name))
}

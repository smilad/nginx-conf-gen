// Code generated by MockGen. DO NOT EDIT.
// Source: nginx/repository_contract (interfaces: IGenerator)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "nginx/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIGenerator is a mock of IGenerator interface.
type MockIGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockIGeneratorMockRecorder
}

// MockIGeneratorMockRecorder is the mock recorder for MockIGenerator.
type MockIGeneratorMockRecorder struct {
	mock *MockIGenerator
}

// NewMockIGenerator creates a new mock instance.
func NewMockIGenerator(ctrl *gomock.Controller) *MockIGenerator {
	mock := &MockIGenerator{ctrl: ctrl}
	mock.recorder = &MockIGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGenerator) EXPECT() *MockIGeneratorMockRecorder {
	return m.recorder
}

// DeleteConfig mocks base method.
func (m *MockIGenerator) DeleteConfig(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConfig indicates an expected call of DeleteConfig.
func (mr *MockIGeneratorMockRecorder) DeleteConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConfig", reflect.TypeOf((*MockIGenerator)(nil).DeleteConfig), arg0, arg1)
}

// GenerateConfig mocks base method.
func (m *MockIGenerator) GenerateConfig(arg0 context.Context, arg1 models.DomainAddr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateConfig indicates an expected call of GenerateConfig.
func (mr *MockIGeneratorMockRecorder) GenerateConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateConfig", reflect.TypeOf((*MockIGenerator)(nil).GenerateConfig), arg0, arg1)
}

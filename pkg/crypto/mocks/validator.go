// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/crypto/validator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTlsValidator is a mock of TlsValidator interface.
type MockTlsValidator struct {
	ctrl     *gomock.Controller
	recorder *MockTlsValidatorMockRecorder
}

// MockTlsValidatorMockRecorder is the mock recorder for MockTlsValidator.
type MockTlsValidatorMockRecorder struct {
	mock *MockTlsValidator
}

// NewMockTlsValidator creates a new mock instance.
func NewMockTlsValidator(ctrl *gomock.Controller) *MockTlsValidator {
	mock := &MockTlsValidator{ctrl: ctrl}
	mock.recorder = &MockTlsValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTlsValidator) EXPECT() *MockTlsValidatorMockRecorder {
	return m.recorder
}

// IsSignedByUnknownAuthority mocks base method.
func (m *MockTlsValidator) IsSignedByUnknownAuthority(host, port string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSignedByUnknownAuthority", host, port)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSignedByUnknownAuthority indicates an expected call of IsSignedByUnknownAuthority.
func (mr *MockTlsValidatorMockRecorder) IsSignedByUnknownAuthority(host, port interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSignedByUnknownAuthority", reflect.TypeOf((*MockTlsValidator)(nil).IsSignedByUnknownAuthority), host, port)
}

// ValidateCert mocks base method.
func (m *MockTlsValidator) ValidateCert(host, port, caCertContent string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCert", host, port, caCertContent)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateCert indicates an expected call of ValidateCert.
func (mr *MockTlsValidatorMockRecorder) ValidateCert(host, port, caCertContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCert", reflect.TypeOf((*MockTlsValidator)(nil).ValidateCert), host, port, caCertContent)
}
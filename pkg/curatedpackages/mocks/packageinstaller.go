// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/curatedpackages/packageinstaller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPackageController is a mock of PackageController interface.
type MockPackageController struct {
	ctrl     *gomock.Controller
	recorder *MockPackageControllerMockRecorder
}

// MockPackageControllerMockRecorder is the mock recorder for MockPackageController.
type MockPackageControllerMockRecorder struct {
	mock *MockPackageController
}

// NewMockPackageController creates a new mock instance.
func NewMockPackageController(ctrl *gomock.Controller) *MockPackageController {
	mock := &MockPackageController{ctrl: ctrl}
	mock.recorder = &MockPackageControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPackageController) EXPECT() *MockPackageControllerMockRecorder {
	return m.recorder
}

// EnableCuratedPackages mocks base method.
func (m *MockPackageController) EnableCuratedPackages(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableCuratedPackages", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableCuratedPackages indicates an expected call of EnableCuratedPackages.
func (mr *MockPackageControllerMockRecorder) EnableCuratedPackages(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableCuratedPackages", reflect.TypeOf((*MockPackageController)(nil).EnableCuratedPackages), ctx)
}

// IsInstalled mocks base method.
func (m *MockPackageController) IsInstalled(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInstalled", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsInstalled indicates an expected call of IsInstalled.
func (mr *MockPackageControllerMockRecorder) IsInstalled(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInstalled", reflect.TypeOf((*MockPackageController)(nil).IsInstalled), ctx)
}

// MockPackageHandler is a mock of PackageHandler interface.
type MockPackageHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPackageHandlerMockRecorder
}

// MockPackageHandlerMockRecorder is the mock recorder for MockPackageHandler.
type MockPackageHandlerMockRecorder struct {
	mock *MockPackageHandler
}

// NewMockPackageHandler creates a new mock instance.
func NewMockPackageHandler(ctrl *gomock.Controller) *MockPackageHandler {
	mock := &MockPackageHandler{ctrl: ctrl}
	mock.recorder = &MockPackageHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPackageHandler) EXPECT() *MockPackageHandlerMockRecorder {
	return m.recorder
}

// CreatePackages mocks base method.
func (m *MockPackageHandler) CreatePackages(ctx context.Context, fileName, kubeConfig string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePackages", ctx, fileName, kubeConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePackages indicates an expected call of CreatePackages.
func (mr *MockPackageHandlerMockRecorder) CreatePackages(ctx, fileName, kubeConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePackages", reflect.TypeOf((*MockPackageHandler)(nil).CreatePackages), ctx, fileName, kubeConfig)
}

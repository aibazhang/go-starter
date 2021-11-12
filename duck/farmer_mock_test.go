// Code generated by MockGen. DO NOT EDIT.
// Source: farmer.go

// Package duck is a generated GoMock package.
package duck

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFarmerI is a mock of FarmerI interface.
type MockFarmerI struct {
	ctrl     *gomock.Controller
	recorder *MockFarmerIMockRecorder
}

// MockFarmerIMockRecorder is the mock recorder for MockFarmerI.
type MockFarmerIMockRecorder struct {
	mock *MockFarmerI
}

// NewMockFarmerI creates a new mock instance.
func NewMockFarmerI(ctrl *gomock.Controller) *MockFarmerI {
	mock := &MockFarmerI{ctrl: ctrl}
	mock.recorder = &MockFarmerIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFarmerI) EXPECT() *MockFarmerIMockRecorder {
	return m.recorder
}

// Breed mocks base method.
func (m *MockFarmerI) Breed() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Breed")
	ret0, _ := ret[0].(string)
	return ret0
}

// Breed indicates an expected call of Breed.
func (mr *MockFarmerIMockRecorder) Breed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Breed", reflect.TypeOf((*MockFarmerI)(nil).Breed))
}

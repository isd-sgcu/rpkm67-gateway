// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/validator/validator.go

// Package mock_validator is a generated GoMock package.
package mock_validator

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDtoValidator is a mock of DtoValidator interface.
type MockDtoValidator struct {
	ctrl     *gomock.Controller
	recorder *MockDtoValidatorMockRecorder
}

// MockDtoValidatorMockRecorder is the mock recorder for MockDtoValidator.
type MockDtoValidatorMockRecorder struct {
	mock *MockDtoValidator
}

// NewMockDtoValidator creates a new mock instance.
func NewMockDtoValidator(ctrl *gomock.Controller) *MockDtoValidator {
	mock := &MockDtoValidator{ctrl: ctrl}
	mock.recorder = &MockDtoValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDtoValidator) EXPECT() *MockDtoValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockDtoValidator) Validate(arg0 interface{}) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockDtoValidatorMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockDtoValidator)(nil).Validate), arg0)
}

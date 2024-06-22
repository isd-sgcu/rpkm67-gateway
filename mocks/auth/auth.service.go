// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/auth/auth.service.go

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	apperrors "github.com/isd-sgcu/rpkm67-gateway/apperrors"
	dto "github.com/isd-sgcu/rpkm67-gateway/internal/dto"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// ForgotPassword mocks base method.
func (m *MockService) ForgotPassword(req *dto.ForgotPasswordRequest) (*dto.ForgotPasswordResponse, *apperrors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotPassword", req)
	ret0, _ := ret[0].(*dto.ForgotPasswordResponse)
	ret1, _ := ret[1].(*apperrors.AppError)
	return ret0, ret1
}

// ForgotPassword indicates an expected call of ForgotPassword.
func (mr *MockServiceMockRecorder) ForgotPassword(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotPassword", reflect.TypeOf((*MockService)(nil).ForgotPassword), req)
}

// RefreshToken mocks base method.
func (m *MockService) RefreshToken() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RefreshToken")
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockServiceMockRecorder) RefreshToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockService)(nil).RefreshToken))
}

// ResetPassword mocks base method.
func (m *MockService) ResetPassword(req *dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, *apperrors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPassword", req)
	ret0, _ := ret[0].(*dto.ResetPasswordResponse)
	ret1, _ := ret[1].(*apperrors.AppError)
	return ret0, ret1
}

// ResetPassword indicates an expected call of ResetPassword.
func (mr *MockServiceMockRecorder) ResetPassword(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockService)(nil).ResetPassword), req)
}

// SignIn mocks base method.
func (m *MockService) SignIn(req *dto.SignInRequest) (*dto.Credential, *apperrors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", req)
	ret0, _ := ret[0].(*dto.Credential)
	ret1, _ := ret[1].(*apperrors.AppError)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockServiceMockRecorder) SignIn(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockService)(nil).SignIn), req)
}

// SignOut mocks base method.
func (m *MockService) SignOut(req *dto.TokenPayloadAuth) (*dto.SignOutResponse, *apperrors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignOut", req)
	ret0, _ := ret[0].(*dto.SignOutResponse)
	ret1, _ := ret[1].(*apperrors.AppError)
	return ret0, ret1
}

// SignOut indicates an expected call of SignOut.
func (mr *MockServiceMockRecorder) SignOut(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignOut", reflect.TypeOf((*MockService)(nil).SignOut), req)
}

// SignUp mocks base method.
func (m *MockService) SignUp(req *dto.SignUpRequest) (*dto.Credential, *apperrors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", req)
	ret0, _ := ret[0].(*dto.Credential)
	ret1, _ := ret[1].(*apperrors.AppError)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockServiceMockRecorder) SignUp(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockService)(nil).SignUp), req)
}

// Validate mocks base method.
func (m *MockService) Validate() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Validate")
}

// Validate indicates an expected call of Validate.
func (mr *MockServiceMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockService)(nil).Validate))
}
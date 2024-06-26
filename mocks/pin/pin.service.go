// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pin/pin.service.go

// Package mock_pin is a generated GoMock package.
package mock_pin

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	apperror "github.com/isd-sgcu/rpkm67-gateway/apperror"
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

// CheckPin mocks base method.
func (m *MockService) CheckPin(req *dto.CheckPinRequest) (*dto.CheckPinResponse, *apperror.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPin", req)
	ret0, _ := ret[0].(*dto.CheckPinResponse)
	ret1, _ := ret[1].(*apperror.AppError)
	return ret0, ret1
}

// CheckPin indicates an expected call of CheckPin.
func (mr *MockServiceMockRecorder) CheckPin(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPin", reflect.TypeOf((*MockService)(nil).CheckPin), req)
}

// FindAll mocks base method.
func (m *MockService) FindAll(req *dto.FindAllPinRequest) (*dto.FindAllPinResponse, *apperror.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", req)
	ret0, _ := ret[0].(*dto.FindAllPinResponse)
	ret1, _ := ret[1].(*apperror.AppError)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockServiceMockRecorder) FindAll(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockService)(nil).FindAll), req)
}

// ResetPin mocks base method.
func (m *MockService) ResetPin(req *dto.ResetPinRequest) (*dto.ResetPinResponse, *apperror.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPin", req)
	ret0, _ := ret[0].(*dto.ResetPinResponse)
	ret1, _ := ret[1].(*apperror.AppError)
	return ret0, ret1
}

// ResetPin indicates an expected call of ResetPin.
func (mr *MockServiceMockRecorder) ResetPin(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPin", reflect.TypeOf((*MockService)(nil).ResetPin), req)
}

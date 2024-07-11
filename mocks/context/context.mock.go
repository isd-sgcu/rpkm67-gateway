// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/context/context.go

// Package mock_context is a generated GoMock package.
package mock_context

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	apperror "github.com/isd-sgcu/rpkm67-gateway/apperror"
	dto "github.com/isd-sgcu/rpkm67-gateway/internal/dto"
	trace "go.opentelemetry.io/otel/trace"
)

// MockCtx is a mock of Ctx interface.
type MockCtx struct {
	ctrl     *gomock.Controller
	recorder *MockCtxMockRecorder
}

// MockCtxMockRecorder is the mock recorder for MockCtx.
type MockCtxMockRecorder struct {
	mock *MockCtx
}

// NewMockCtx creates a new mock instance.
func NewMockCtx(ctrl *gomock.Controller) *MockCtx {
	mock := &MockCtx{ctrl: ctrl}
	mock.recorder = &MockCtxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCtx) EXPECT() *MockCtxMockRecorder {
	return m.recorder
}

// BadRequestError mocks base method.
func (m *MockCtx) BadRequestError(err string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BadRequestError", err)
}

// BadRequestError indicates an expected call of BadRequestError.
func (mr *MockCtxMockRecorder) BadRequestError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BadRequestError", reflect.TypeOf((*MockCtx)(nil).BadRequestError), err)
}

// Bind mocks base method.
func (m *MockCtx) Bind(obj interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind.
func (mr *MockCtxMockRecorder) Bind(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockCtx)(nil).Bind), obj)
}

// FormFile mocks base method.
func (m *MockCtx) FormFile(key string, allowedContentType map[string]struct{}, maxFileSize int64) (*dto.DecomposedFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormFile", key, allowedContentType, maxFileSize)
	ret0, _ := ret[0].(*dto.DecomposedFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormFile indicates an expected call of FormFile.
func (mr *MockCtxMockRecorder) FormFile(key, allowedContentType, maxFileSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormFile", reflect.TypeOf((*MockCtx)(nil).FormFile), key, allowedContentType, maxFileSize)
}

// GetHeader mocks base method.
func (m *MockCtx) GetHeader(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHeader indicates an expected call of GetHeader.
func (mr *MockCtxMockRecorder) GetHeader(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockCtx)(nil).GetHeader), key)
}

// GetString mocks base method.
func (m *MockCtx) GetString(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetString", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetString indicates an expected call of GetString.
func (mr *MockCtxMockRecorder) GetString(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockCtx)(nil).GetString), key)
}

// GetTracer mocks base method.
func (m *MockCtx) GetTracer() trace.Tracer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracer")
	ret0, _ := ret[0].(trace.Tracer)
	return ret0
}

// GetTracer indicates an expected call of GetTracer.
func (mr *MockCtxMockRecorder) GetTracer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracer", reflect.TypeOf((*MockCtx)(nil).GetTracer))
}

// InternalServerError mocks base method.
func (m *MockCtx) InternalServerError(err string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InternalServerError", err)
}

// InternalServerError indicates an expected call of InternalServerError.
func (mr *MockCtxMockRecorder) InternalServerError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InternalServerError", reflect.TypeOf((*MockCtx)(nil).InternalServerError), err)
}

// JSON mocks base method.
func (m *MockCtx) JSON(statusCode int, obj interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "JSON", statusCode, obj)
}

// JSON indicates an expected call of JSON.
func (mr *MockCtxMockRecorder) JSON(statusCode, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSON", reflect.TypeOf((*MockCtx)(nil).JSON), statusCode, obj)
}

// NewUUID mocks base method.
func (m *MockCtx) NewUUID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUUID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// NewUUID indicates an expected call of NewUUID.
func (mr *MockCtxMockRecorder) NewUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUUID", reflect.TypeOf((*MockCtx)(nil).NewUUID))
}

// Param mocks base method.
func (m *MockCtx) Param(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Param", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// Param indicates an expected call of Param.
func (mr *MockCtxMockRecorder) Param(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param", reflect.TypeOf((*MockCtx)(nil).Param), key)
}

// PostForm mocks base method.
func (m *MockCtx) PostForm(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostForm", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// PostForm indicates an expected call of PostForm.
func (mr *MockCtxMockRecorder) PostForm(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostForm", reflect.TypeOf((*MockCtx)(nil).PostForm), key)
}

// Query mocks base method.
func (m *MockCtx) Query(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockCtxMockRecorder) Query(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockCtx)(nil).Query), key)
}

// RequestContext mocks base method.
func (m *MockCtx) RequestContext() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestContext")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// RequestContext indicates an expected call of RequestContext.
func (mr *MockCtxMockRecorder) RequestContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestContext", reflect.TypeOf((*MockCtx)(nil).RequestContext))
}

// ResponseError mocks base method.
func (m *MockCtx) ResponseError(err *apperror.AppError) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResponseError", err)
}

// ResponseError indicates an expected call of ResponseError.
func (mr *MockCtxMockRecorder) ResponseError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseError", reflect.TypeOf((*MockCtx)(nil).ResponseError), err)
}

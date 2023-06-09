// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock is a generated GoMock package.
package mock

import (
	reportstore "src/reportstore"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateReport mocks base method
func (m *MockStore) CreateReport(r reportstore.CreateReportRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReport", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReport indicates an expected call of CreateReport
func (mr *MockStoreMockRecorder) CreateReport(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReport", reflect.TypeOf((*MockStore)(nil).CreateReport), r)
}

// UpdateReportStatus mocks base method
func (m *MockStore) UpdateReportStatus(r reportstore.UpdateReportStatusRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReportStatus", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReportStatus indicates an expected call of UpdateReportStatus
func (mr *MockStoreMockRecorder) UpdateReportStatus(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReportStatus", reflect.TypeOf((*MockStore)(nil).UpdateReportStatus), r)
}

// GetUserReport mocks base method
func (m *MockStore) GetUserReport(r reportstore.GetUserReportRequest) (reportstore.GetUserReportResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserReport", r)
	ret0, _ := ret[0].(reportstore.GetUserReportResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserReport indicates an expected call of GetUserReport
func (mr *MockStoreMockRecorder) GetUserReport(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserReport", reflect.TypeOf((*MockStore)(nil).GetUserReport), r)
}

// GetUserReports mocks base method
func (m *MockStore) GetUserReports(r reportstore.GetUserReportsRequest) (reportstore.GetUserReportsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserReports", r)
	ret0, _ := ret[0].(reportstore.GetUserReportsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserReports indicates an expected call of GetUserReports
func (mr *MockStoreMockRecorder) GetUserReports(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserReports", reflect.TypeOf((*MockStore)(nil).GetUserReports), r)
}

// UserHasAccess mocks base method
func (m *MockStore) UserHasAccess(r reportstore.UserHasAccessRequest) (reportstore.UserHasAccessResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserHasAccess", r)
	ret0, _ := ret[0].(reportstore.UserHasAccessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserHasAccess indicates an expected call of UserHasAccess
func (mr *MockStoreMockRecorder) UserHasAccess(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserHasAccess", reflect.TypeOf((*MockStore)(nil).UserHasAccess), r)
}

// AppendMessage mocks base method
func (m *MockStore) AppendMessage(r reportstore.AppendMessageRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendMessage", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendMessage indicates an expected call of AppendMessage
func (mr *MockStoreMockRecorder) AppendMessage(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendMessage", reflect.TypeOf((*MockStore)(nil).AppendMessage), r)
}
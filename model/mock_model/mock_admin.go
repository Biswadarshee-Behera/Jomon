// Code generated by MockGen. DO NOT EDIT.
// Source: admin.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	model "github.com/traPtitech/Jomon/model"
	reflect "reflect"
)

// MockAdminRepository is a mock of AdminRepository interface
type MockAdminRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAdminRepositoryMockRecorder
}

// MockAdminRepositoryMockRecorder is the mock recorder for MockAdminRepository
type MockAdminRepositoryMockRecorder struct {
	mock *MockAdminRepository
}

// NewMockAdminRepository creates a new mock instance
func NewMockAdminRepository(ctrl *gomock.Controller) *MockAdminRepository {
	mock := &MockAdminRepository{ctrl: ctrl}
	mock.recorder = &MockAdminRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdminRepository) EXPECT() *MockAdminRepositoryMockRecorder {
	return m.recorder
}

// GetAdmins mocks base method
func (m *MockAdminRepository) GetAdmins(ctx context.Context) ([]*model.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmins", ctx)
	ret0, _ := ret[0].([]*model.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdmins indicates an expected call of GetAdmins
func (mr *MockAdminRepositoryMockRecorder) GetAdmins(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmins", reflect.TypeOf((*MockAdminRepository)(nil).GetAdmins), ctx)
}

// CreateAdmin mocks base method
func (m *MockAdminRepository) CreateAdmin(ctx context.Context, userID uuid.UUID) (*model.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdmin", ctx, userID)
	ret0, _ := ret[0].(*model.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdmin indicates an expected call of CreateAdmin
func (mr *MockAdminRepositoryMockRecorder) CreateAdmin(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockAdminRepository)(nil).CreateAdmin), ctx, userID)
}

// DeleteAdmin mocks base method
func (m *MockAdminRepository) DeleteAdmin(ctx context.Context, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdmin", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAdmin indicates an expected call of DeleteAdmin
func (mr *MockAdminRepositoryMockRecorder) DeleteAdmin(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdmin", reflect.TypeOf((*MockAdminRepository)(nil).DeleteAdmin), ctx, userID)
}

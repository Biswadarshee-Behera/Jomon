// Code generated by MockGen. DO NOT EDIT.
// Source: request.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/traPtitech/Jomon/model"
	reflect "reflect"
)

// MockRequestRepository is a mock of RequestRepository interface
type MockRequestRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRequestRepositoryMockRecorder
}

// MockRequestRepositoryMockRecorder is the mock recorder for MockRequestRepository
type MockRequestRepositoryMockRecorder struct {
	mock *MockRequestRepository
}

// NewMockRequestRepository creates a new mock instance
func NewMockRequestRepository(ctrl *gomock.Controller) *MockRequestRepository {
	mock := &MockRequestRepository{ctrl: ctrl}
	mock.recorder = &MockRequestRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRequestRepository) EXPECT() *MockRequestRepositoryMockRecorder {
	return m.recorder
}

// GetRequests mocks base method
func (m *MockRequestRepository) GetRequests(ctx context.Context, query model.RequestQuery) ([]*model.RequestResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequests", ctx, query)
	ret0, _ := ret[0].([]*model.RequestResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequests indicates an expected call of GetRequests
func (mr *MockRequestRepositoryMockRecorder) GetRequests(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequests", reflect.TypeOf((*MockRequestRepository)(nil).GetRequests), ctx, query)
}

// CreateRequest mocks base method
func (m *MockRequestRepository) CreateRequest(ctx context.Context, amount int, title, content string, tags []*model.Tag, group model.Group, files []*model.File) (*model.RequestDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRequest", ctx, amount, title, content, tags, group, files)
	ret0, _ := ret[0].(*model.RequestDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRequest indicates an expected call of CreateRequest
func (mr *MockRequestRepositoryMockRecorder) CreateRequest(ctx, amount, title, content, tags, group, files interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRequest", reflect.TypeOf((*MockRequestRepository)(nil).CreateRequest), ctx, amount, title, content, tags, group, files)
}
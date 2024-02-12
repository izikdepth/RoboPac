// Code generated by MockGen. DO NOT EDIT.
// Source: ./turboswap/interface.go
//
// Generated by this command:
//
//	mockgen -source=./turboswap/interface.go -destination=./turboswap/mock.go -package=turboswap
//

// Package turboswap is a generated GoMock package.
package turboswap

import (
	context "context"
	reflect "reflect"

	store "github.com/kehiy/RoboPac/store"
	gomock "go.uber.org/mock/gomock"
)

// MockITurboSwap is a mock of ITurboSwap interface.
type MockITurboSwap struct {
	ctrl     *gomock.Controller
	recorder *MockITurboSwapMockRecorder
}

// MockITurboSwapMockRecorder is the mock recorder for MockITurboSwap.
type MockITurboSwapMockRecorder struct {
	mock *MockITurboSwap
}

// NewMockITurboSwap creates a new mock instance.
func NewMockITurboSwap(ctrl *gomock.Controller) *MockITurboSwap {
	mock := &MockITurboSwap{ctrl: ctrl}
	mock.recorder = &MockITurboSwapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITurboSwap) EXPECT() *MockITurboSwapMockRecorder {
	return m.recorder
}

// GetStatus mocks base method.
func (m *MockITurboSwap) GetStatus(ctx context.Context, party *store.TwitterParty) (*DiscountStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus", ctx, party)
	ret0, _ := ret[0].(*DiscountStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockITurboSwapMockRecorder) GetStatus(ctx, party any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockITurboSwap)(nil).GetStatus), ctx, party)
}

// SendDiscountCode mocks base method.
func (m *MockITurboSwap) SendDiscountCode(ctx context.Context, party *store.TwitterParty) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDiscountCode", ctx, party)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendDiscountCode indicates an expected call of SendDiscountCode.
func (mr *MockITurboSwapMockRecorder) SendDiscountCode(ctx, party any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDiscountCode", reflect.TypeOf((*MockITurboSwap)(nil).SendDiscountCode), ctx, party)
}

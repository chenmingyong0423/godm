// Code generated by MockGen. DO NOT EDIT.
// Source: finder.go
//
// Generated by this command:
//
//	mockgen -source=finder.go -destination=../mock/finder.mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	options "go.mongodb.org/mongo-driver/mongo/options"
	gomock "go.uber.org/mock/gomock"
)

// MockiFinder is a mock of iFinder interface.
type MockiFinder[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockiFinderMockRecorder[T]
}

// MockiFinderMockRecorder is the mock recorder for MockiFinder.
type MockiFinderMockRecorder[T any] struct {
	mock *MockiFinder[T]
}

// NewMockiFinder creates a new mock instance.
func NewMockiFinder[T any](ctrl *gomock.Controller) *MockiFinder[T] {
	mock := &MockiFinder[T]{ctrl: ctrl}
	mock.recorder = &MockiFinderMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiFinder[T]) EXPECT() *MockiFinderMockRecorder[T] {
	return m.recorder
}

// Count mocks base method.
func (m *MockiFinder[T]) Count(ctx context.Context, opts ...*options.CountOptions) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Count", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockiFinderMockRecorder[T]) Count(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockiFinder[T])(nil).Count), varargs...)
}

// Find mocks base method.
func (m *MockiFinder[T]) Find(ctx context.Context, opts ...*options.FindOptions) ([]*T, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].([]*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockiFinderMockRecorder[T]) Find(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockiFinder[T])(nil).Find), varargs...)
}

// FindOne mocks base method.
func (m *MockiFinder[T]) FindOne(ctx context.Context, opts ...*options.FindOneOptions) (*T, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockiFinderMockRecorder[T]) FindOne(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockiFinder[T])(nil).FindOne), varargs...)
}

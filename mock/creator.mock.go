// Code generated by MockGen. DO NOT EDIT.
// Source: creator.go
//
// Generated by this command:
//
//	mockgen -source=creator.go -destination=../mock/creator.mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	mongo "go.mongodb.org/mongo-driver/v2/mongo"
	options "go.mongodb.org/mongo-driver/v2/mongo/options"
	gomock "go.uber.org/mock/gomock"
)

// MockiCreator is a mock of iCreator interface.
type MockiCreator[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockiCreatorMockRecorder[T]
}

// MockiCreatorMockRecorder is the mock recorder for MockiCreator.
type MockiCreatorMockRecorder[T any] struct {
	mock *MockiCreator[T]
}

// NewMockiCreator creates a new mock instance.
func NewMockiCreator[T any](ctrl *gomock.Controller) *MockiCreator[T] {
	mock := &MockiCreator[T]{ctrl: ctrl}
	mock.recorder = &MockiCreatorMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiCreator[T]) EXPECT() *MockiCreatorMockRecorder[T] {
	return m.recorder
}

// InsertMany mocks base method.
func (m *MockiCreator[T]) InsertMany(ctx context.Context, docs []*T, opts ...options.Lister[options.InsertManyOptions]) (*mongo.InsertManyResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, docs}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertMany", varargs...)
	ret0, _ := ret[0].(*mongo.InsertManyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMany indicates an expected call of InsertMany.
func (mr *MockiCreatorMockRecorder[T]) InsertMany(ctx, docs any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, docs}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMany", reflect.TypeOf((*MockiCreator[T])(nil).InsertMany), varargs...)
}

// InsertOne mocks base method.
func (m *MockiCreator[T]) InsertOne(ctx context.Context, docs *T, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, docs}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertOne", varargs...)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockiCreatorMockRecorder[T]) InsertOne(ctx, docs any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, docs}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockiCreator[T])(nil).InsertOne), varargs...)
}

// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deleter

import (
	"context"
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/field"

	"github.com/chenmingyong0423/go-mongox/v2/callback"

	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//go:generate mockgen -source=deleter.go -destination=../mock/deleter.mock.go -package=mocks
type IDeleter[T any] interface {
	DeleteOne(ctx context.Context, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, opts ...options.Lister[options.DeleteManyOptions]) (*mongo.DeleteResult, error)
}

var _ IDeleter[any] = (*Deleter[any])(nil)

func NewDeleter[T any](collection *mongo.Collection, dbCallbacks *callback.Callback, fields []*field.Filed) *Deleter[T] {
	return &Deleter[T]{collection: collection, dbCallbacks: dbCallbacks, fields: fields}
}

type Deleter[T any] struct {
	collection *mongo.Collection
	fields     []*field.Filed

	filter    any
	modelHook any

	dbCallbacks *callback.Callback
	beforeHooks []beforeHookFn
	afterHooks  []afterHookFn
}

func (d *Deleter[T]) RegisterBeforeHooks(hooks ...beforeHookFn) *Deleter[T] {
	d.beforeHooks = append(d.beforeHooks, hooks...)
	return d
}

func (d *Deleter[T]) RegisterAfterHooks(hooks ...afterHookFn) *Deleter[T] {
	d.afterHooks = append(d.afterHooks, hooks...)
	return d
}

func (d *Deleter[T]) preActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error {
	err := d.dbCallbacks.Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, beforeHook := range d.beforeHooks {
		err = beforeHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Deleter[T]) postActionHandler(ctx context.Context, globalOpContext *operation.OpContext, opContext *OpContext, opType operation.OpType) error {
	err := d.dbCallbacks.Execute(ctx, globalOpContext, opType)
	if err != nil {
		return err
	}
	for _, afterHook := range d.afterHooks {
		err = afterHook(ctx, opContext)
		if err != nil {
			return err
		}
	}
	return nil
}

// Filter is used to set the filter of the query
func (d *Deleter[T]) Filter(filter any) *Deleter[T] {
	d.filter = filter
	return d
}

func (d *Deleter[T]) ModelHook(modelHook any) *Deleter[T] {
	d.modelHook = modelHook
	return d
}

func (d *Deleter[T]) DeleteOne(ctx context.Context, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
	currentTime := time.Now()
	globalOpContext := operation.NewOpContext(d.collection, operation.WithFilter(d.filter), operation.WithMongoOptions(opts), operation.WithModelHook(d.modelHook), operation.WithFields(d.fields), operation.WithStartTime(currentTime))
	opContext := NewOpContext(d.collection, d.filter, WithMongoOptions(opts), WithModelHook(d.modelHook), WithFields(d.fields), WithStartTime(currentTime))
	err := d.preActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeDelete)
	if err != nil {
		return nil, err
	}

	result, err := d.collection.DeleteOne(ctx, d.filter, opts...)
	if err != nil {
		return nil, err
	}

	globalOpContext.Result = result
	opContext.Result = result
	err = d.postActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterDelete)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *Deleter[T]) DeleteMany(ctx context.Context, opts ...options.Lister[options.DeleteManyOptions]) (*mongo.DeleteResult, error) {
	currentTime := time.Now()
	globalOpContext := operation.NewOpContext(d.collection, operation.WithFilter(d.filter), operation.WithMongoOptions(opts), operation.WithModelHook(d.modelHook), operation.WithFields(d.fields), operation.WithStartTime(currentTime))
	opContext := NewOpContext(d.collection, d.filter, WithMongoOptions(opts), WithModelHook(d.modelHook), WithFields(d.fields), WithStartTime(currentTime))
	err := d.preActionHandler(ctx, globalOpContext, opContext, operation.OpTypeBeforeDelete)
	if err != nil {
		return nil, err
	}

	result, err := d.collection.DeleteMany(ctx, d.filter, opts...)
	if err != nil {
		return nil, err
	}

	globalOpContext.Result = result
	opContext.Result = result
	err = d.postActionHandler(ctx, globalOpContext, opContext, operation.OpTypeAfterDelete)
	if err != nil {
		return nil, err
	}

	return result, nil
}

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

//go:build e2e

package mongox

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type testUser struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	Age          int
	UnknownField string `bson:"-"`
}

func TestCollection_e2e_FindOne(t *testing.T) {
	collection := getCollection(t)
	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx    context.Context
		filter any
		opts   []*options.FindOneOptions

		wantT   *testUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "not map, bson.D, struct and struct pointer",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:    context.Background(),
			filter: 1,

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name:   "nil filter",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:    context.Background(),
			filter: nil,

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "empty bson.D filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: bson.D{},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by bson.D filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: bson.D{bson.E{Key: id, Value: "123"}},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "empty map filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: map[string]any{},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by map filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx: context.Background(),
			filter: map[string]any{
				"_id": "123",
			},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "zero struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: testUser{},

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if !errors.Is(err, mongo.ErrNoDocuments) {
					t.Errorf("expected an error but not eq: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: testUser{Id: "123", Name: "cmy", Age: 18},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "zero struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: &testUser{},

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if !errors.Is(err, mongo.ErrNoDocuments) {
					t.Errorf("expected an error but not eq: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: &testUser{Id: "123", Name: "cmy", Age: 18},

			wantT: &testUser{Id: "123", Name: "cmy", Age: 18},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
					return false
				}
				return true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			gotT, err := collection.FindOne(tc.ctx, tc.filter, tc.opts...)
			tc.after(tc.ctx, t)
			assert.True(t, tc.wantErr(t, err))
			assert.Equal(t, tc.wantT, gotT)
		})
	}
}

func TestCollection_e2e_Find(t *testing.T) {
	collection := getCollection(t)
	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx    context.Context
		filter any
		opts   []*options.FindOptions

		wantT   []*testUser
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "not map, bson.D, struct and struct pointer",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:    context.Background(),
			filter: 1,

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "empty bson.D filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: bson.D{},

			wantT: []*testUser{
				{Id: "123", Name: "cmy", Age: 18},
				{Id: "456", Name: "cmy", Age: 18},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "decode failed",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, NewBsonBuilder().Add(id, "123").Add("name", "cmy").Add("age", "18").Build())
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: NewBsonBuilder().Id("123").Build(),

			wantT: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err == nil {
					t.Errorf("expected an error but got none")
					return false
				}
				return true
			},
		},
		{
			name: "get one by bson.D filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: bson.D{bson.E{Key: id, Value: "123"}},

			wantT: []*testUser{
				{Id: "123", Name: "cmy", Age: 18},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "empty map filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: map[string]any{},

			wantT: []*testUser{
				{Id: "123", Name: "cmy", Age: 18},
				{Id: "456", Name: "cmy", Age: 18},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by map filter",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx: context.Background(),
			filter: map[string]any{
				"_id": "123",
			},

			wantT: []*testUser{
				{Id: "123", Name: "cmy", Age: 18},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "zero struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: testUser{},

			wantT: []*testUser{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by struct",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: testUser{Id: "123", Name: "cmy", Age: 18},

			wantT: []*testUser{
				{Id: "123", Name: "cmy", Age: 18},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "zero struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: &testUser{},

			wantT: []*testUser{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
		{
			name: "get one by struct pointer",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
				_, fErr = collection.collection.InsertOne(ctx, testData{
					Id:   "456",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)

			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
				_, fErr = collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("456").Build())
				assert.NoError(t, fErr)
			},

			ctx:    context.Background(),
			filter: &testUser{Id: "123", Name: "cmy", Age: 18},

			wantT: []*testUser{
				{Id: "123", Name: "cmy", Age: 18},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
					return false
				}
				return true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			gotT, err := collection.Find(tc.ctx, tc.filter, tc.opts...)
			tc.after(tc.ctx, t)
			assert.True(t, tc.wantErr(t, err))
			assert.Equal(t, tc.wantT, gotT)
		})
	}
}

func TestCollection_e2e_FindById(t *testing.T) {
	collection := getCollection(t)

	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx  context.Context
		id   string
		opts []*options.FindOneOptions

		wantT   *testUser
		wantErr error
	}{
		{
			name:   "no document",
			before: func(_ context.Context, _ *testing.T) {},
			after:  func(_ context.Context, _ *testing.T) {},

			ctx:  context.Background(),
			id:   "123",
			opts: nil,

			wantT:   nil,
			wantErr: mongo.ErrNoDocuments,
		},
		{
			name: "found",
			before: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.InsertOne(ctx, testData{
					Id:   "123",
					Name: "cmy",
					Age:  18,
				})
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := collection.collection.DeleteOne(ctx, NewBsonBuilder().Id("123").Build())
				assert.NoError(t, fErr)
			},

			ctx:  context.Background(),
			id:   "123",
			opts: nil,

			wantT: &testUser{
				Id:   "123",
				Name: "cmy",
				Age:  18,
			},
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			gotT, err := collection.FindById(tc.ctx, tc.id, tc.opts...)
			tc.after(tc.ctx, t)
			assert.Equal(t, tc.wantT, gotT)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func getCollection(t *testing.T) *Collection[testUser] {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	assert.NoError(t, err)
	assert.NoError(t, client.Ping(context.Background(), readpref.Primary()))

	collection := NewCollection[testUser](client.Database("db-test").Collection("test_user"))
	return collection
}

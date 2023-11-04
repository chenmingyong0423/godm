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

package aggregation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_comparisonBuilder_Eq(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$eq", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$eq", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			expected:    bson.D{bson.E{Key: "$eq", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Eq(tc.expressions...).Build())
		})
	}
}

func Test_comparisonBuilder_Ne(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$ne", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$ne", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			expected:    bson.D{bson.E{Key: "$ne", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Ne(tc.expressions...).Build())
		})
	}
}

func Test_comparisonBuilder_Gt(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$gt", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$gt", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			expected:    bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Gt(tc.expressions...).Build())
		})
	}
}

func Test_comparisonBuilder_Gte(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$gte", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$gte", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			expected:    bson.D{bson.E{Key: "$gte", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Gte(tc.expressions...).Build())
		})
	}
}

func Test_comparisonBuilder_Lt(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$lt", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$lt", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			expected:    bson.D{bson.E{Key: "$lt", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Lt(tc.expressions...).Build())
		})
	}
}

func Test_comparisonBuilder_Lte(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$lte", Value: []any{nil}}},
		},
		{
			name:        "empty",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$lte", Value: []any{}}},
		},
		{
			name:        "normal",
			expressions: []any{"$qty", 250},
			expected:    bson.D{bson.E{Key: "$lte", Value: []any{"$qty", 250}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Lte(tc.expressions...).Build())
		})
	}
}

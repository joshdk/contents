// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package contents

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelpers(t *testing.T) {

	tests := []struct {
		title   string
		ctx     context.Context
		keys    []interface{}
		pairs   []Pair
		mapping map[interface{}]interface{}
	}{
		{
			title:   "nil context",
			ctx:     nil,
			mapping: map[interface{}]interface{}{},
		},
		{
			title:   "background context",
			ctx:     context.Background(),
			mapping: map[interface{}]interface{}{},
		},
		{
			title:   "todo context",
			ctx:     context.TODO(),
			mapping: map[interface{}]interface{}{},
		},
		{
			title: "cancel context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx, cancel := context.WithCancel(ctx)
				_ = cancel
				return ctx
			}(),
			mapping: map[interface{}]interface{}{},
		},
		{
			title: "canceled cancel context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx, cancel := context.WithCancel(ctx)
				cancel()
				return ctx
			}(),
			mapping: map[interface{}]interface{}{},
		},
		{
			title: "timeout context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 0)
				_ = cancel
				return ctx
			}(),
			mapping: map[interface{}]interface{}{},
		},
		{
			title: "canceled timeout context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 0)
				cancel()
				return ctx
			}(),
			mapping: map[interface{}]interface{}{},
		},
		{
			title: "single key context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx = context.WithValue(ctx, "key", "value")
				return ctx
			}(),
			keys: []interface{}{"key"},
			pairs: []Pair{
				{"key", "value"},
			},
			mapping: map[interface{}]interface{}{
				"key": "value",
			},
		},
		{
			title: "multi key context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx = context.WithValue(ctx, "key-1", "value-1")
				ctx = context.WithValue(ctx, "key-2", "value-2")
				ctx = context.WithValue(ctx, "key-3", "value-3")
				return ctx
			}(),
			keys: []interface{}{"key-1", "key-2", "key-3"},
			pairs: []Pair{
				{"key-1", "value-1"},
				{"key-2", "value-2"},
				{"key-3", "value-3"},
			},
			mapping: map[interface{}]interface{}{
				"key-1": "value-1",
				"key-2": "value-2",
				"key-3": "value-3",
			},
		},
		{
			title: "cancel context with keys",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx = context.WithValue(ctx, "key-1", "value-1")
				ctx = context.WithValue(ctx, "key-2", "value-2")
				ctx, cancel := context.WithCancel(ctx)
				_ = cancel
				ctx = context.WithValue(ctx, "key-3", "value-3")
				return ctx
			}(),
			keys: []interface{}{"key-1", "key-2", "key-3"},
			pairs: []Pair{
				{"key-1", "value-1"},
				{"key-2", "value-2"},
				{"key-3", "value-3"},
			},
			mapping: map[interface{}]interface{}{
				"key-1": "value-1",
				"key-2": "value-2",
				"key-3": "value-3",
			},
		},
		{
			title: "timeout context with keys",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx = context.WithValue(ctx, "key-1", "value-1")
				ctx = context.WithValue(ctx, "key-2", "value-2")
				ctx, cancel := context.WithTimeout(ctx, 0)
				_ = cancel
				ctx = context.WithValue(ctx, "key-3", "value-3")
				return ctx
			}(),
			keys: []interface{}{"key-1", "key-2", "key-3"},
			pairs: []Pair{
				{"key-1", "value-1"},
				{"key-2", "value-2"},
				{"key-3", "value-3"},
			},
			mapping: map[interface{}]interface{}{
				"key-1": "value-1",
				"key-2": "value-2",
				"key-3": "value-3",
			},
		},
		{
			title: "duplicate key context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx = context.WithValue(ctx, "key-1", "value-1")
				ctx = context.WithValue(ctx, "key-2", "value-2")
				ctx = context.WithValue(ctx, "key-3", "value-3")
				ctx = context.WithValue(ctx, "key-2", "VALUE-TWO")
				return ctx
			}(),
			keys: []interface{}{"key-1", "key-2", "key-3", "key-2"},
			pairs: []Pair{
				{"key-1", "value-1"},
				{"key-2", "value-2"},
				{"key-3", "value-3"},
				{"key-2", "VALUE-TWO"},
			},
			mapping: map[interface{}]interface{}{
				"key-1": "value-1",
				"key-2": "VALUE-TWO",
				"key-3": "value-3",
			},
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.title)

		t.Run(name, func(t *testing.T) {

			keys := Keys(test.ctx)
			pairs := Pairs(test.ctx)
			mapping := Map(test.ctx)

			require.Equal(t, len(keys), len(pairs))

			for index := range keys {
				assert.Equal(t, keys[index], pairs[index].Key)
			}

			assert.Equal(t, test.keys, keys)
			assert.Equal(t, test.pairs, pairs)
			assert.Equal(t, test.mapping, mapping)

			for key, value := range mapping {
				assert.Equal(t, test.ctx.Value(key), value)
			}

		})

	}

}

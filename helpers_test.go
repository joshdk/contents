// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package contents

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {

	tests := []struct {
		title string
		ctx   context.Context
		keys  []interface{}
	}{
		{
			title: "nil context",
			ctx:   nil,
		},
		{
			title: "background context",
			ctx:   context.Background(),
		},
		{
			title: "todo context",
			ctx:   context.TODO(),
		},
		{
			title: "cancel context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx, cancel := context.WithCancel(ctx)
				_ = cancel
				return ctx
			}(),
		},
		{
			title: "canceled cancel context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx, cancel := context.WithCancel(ctx)
				cancel()
				return ctx
			}(),
		},
		{
			title: "timeout context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 0)
				_ = cancel
				return ctx
			}(),
		},
		{
			title: "canceled timeout context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 0)
				cancel()
				return ctx
			}(),
		},
		{
			title: "single key context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx = context.WithValue(ctx, "key", "value")
				return ctx
			}(),
			keys: []interface{}{"key"},
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
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.title)

		t.Run(name, func(t *testing.T) {

			actual := Keys(test.ctx)

			assert.Equal(t, test.keys, actual)

			for _, key := range actual {
				assert.NotNil(t, test.ctx.Value(key))
			}

		})

	}

}

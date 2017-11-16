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

func TestUnwrap(t *testing.T) {

	tests := []struct {
		title   string
		wrapper func() (context.Context, context.Context)
	}{
		{
			title: "nil context",
			wrapper: func() (context.Context, context.Context) {
				return nil, nil
			},
		},
		{
			title: "background context",
			wrapper: func() (context.Context, context.Context) {
				return nil, context.Background()
			},
		},
		{
			title: "todo context",
			wrapper: func() (context.Context, context.Context) {
				return nil, context.TODO()
			},
		},
		{
			title: "cancel context",
			wrapper: func() (context.Context, context.Context) {
				original := context.Background()
				wrapped, cancel := context.WithCancel(original)
				_ = cancel
				return original, wrapped
			},
		},
		{
			title: "cancel context",
			wrapper: func() (context.Context, context.Context) {
				original := context.Background()
				wrapped, cancel := context.WithTimeout(original, 0)
				_ = cancel
				return original, wrapped
			},
		},
		{
			title: "single key context",
			wrapper: func() (context.Context, context.Context) {
				original := context.Background()
				wrapped := context.WithValue(original, "key", "value")
				return original, wrapped
			},
		},
		{
			title: "multi key context",
			wrapper: func() (context.Context, context.Context) {
				original := context.Background()
				wrapped1 := context.WithValue(original, "key-1", "value-1")
				wrapped2 := context.WithValue(wrapped1, "key-2", "value-2")
				wrapped3 := context.WithValue(wrapped2, "key-3", "value-3")
				return wrapped2, wrapped3
			},
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.title)

		t.Run(name, func(t *testing.T) {

			original, wrapped := test.wrapper()

			unwrapped := Unwrap(wrapped)

			assert.Equal(t, original, unwrapped)

		})

	}

}

func TestKey(t *testing.T) {

	tests := []struct {
		title string
		ctx   context.Context
		key   interface{}
		found bool
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
			key:   "key",
			found: true,
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
			key:   "key-3",
			found: true,
		},
		{
			title: "int key context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx = context.WithValue(ctx, 9001, "value")
				return ctx
			}(),
			key:   9001,
			found: true,
		},
		{
			title: "array key context",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx = context.WithValue(ctx, [1]string{"key"}, "value")
				return ctx
			}(),
			key:   [1]string{"key"},
			found: true,
		},
		{
			title: "cancel context after values",
			ctx: func() context.Context {
				ctx := context.Background()
				ctx = context.WithValue(ctx, "key", "value")
				ctx, cancel := context.WithCancel(ctx)
				_ = cancel
				return ctx
			}(),
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.title)

		t.Run(name, func(t *testing.T) {

			key, found := Key(test.ctx)

			assert.Equal(t, test.found, found)

			assert.Equal(t, test.key, key)

		})

	}

}

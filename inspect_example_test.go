// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package contents_test

import (
	"context"
	"fmt"

	"github.com/joshdk/contents"
)

func ExampleKey_found() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key a", "value a")

	if key, found := contents.Key(ctx); found {
		fmt.Printf("A key was found and its value is %q\n", key)
	} else {
		fmt.Println("No key was found")
	}
	// Output:
	// A key was found and its value is "key a"
}

func ExampleKey_missing() {
	ctx := context.Background()

	if key, found := contents.Key(ctx); found {
		fmt.Printf("A key was found and its value is %q\n", key)
	} else {
		fmt.Println("No key was found")
	}
	// Output:
	// No key was found
}

func ExampleUnwrap_succeeded() {
	ctxBg := context.Background()
	ctxVal := context.WithValue(ctxBg, "key a", "value a")

	ctx := contents.Unwrap(ctxVal)
	if ctx == ctxBg {
		fmt.Println("The context ctx is equal to ctxBg")
	}
	// Output:
	// The context ctx is equal to ctxBg
}

func ExampleUnwrap_failed() {
	ctxBg := context.Background()

	ctx := contents.Unwrap(ctxBg)
	if ctx == nil {
		fmt.Println("The context ctxBg could not be unwrapped")
	}
	// Output:
	// The context ctxBg could not be unwrapped
}

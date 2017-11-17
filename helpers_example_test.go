// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package contents_test

import (
	"context"
	"fmt"
	"sort"

	"github.com/joshdk/contents"
)

func ExampleKeys() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key a", "value a")
	ctx = context.WithValue(ctx, "key b", "value b")
	ctx = context.WithValue(ctx, "key c", "value c")
	ctx = context.WithValue(ctx, "key b", "VALUE B")

	keys := contents.Keys(ctx)

	for index, key := range keys {
		fmt.Printf("Value of keys[%d] is %q\n", index, key)
	}
	// Output:
	// Value of keys[0] is "key a"
	// Value of keys[1] is "key b"
	// Value of keys[2] is "key c"
	// Value of keys[3] is "key b"
}

func ExamplePairs() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key a", "value a")
	ctx = context.WithValue(ctx, "key b", "value b")
	ctx = context.WithValue(ctx, "key c", "value c")
	ctx = context.WithValue(ctx, "key b", "VALUE B")

	pairs := contents.Pairs(ctx)

	for index, pair := range pairs {
		fmt.Printf("Value of pairs[%d] is %q\n", index, pair)
	}
	// Output:
	// Value of pairs[0] is {"key a" "value a"}
	// Value of pairs[1] is {"key b" "value b"}
	// Value of pairs[2] is {"key c" "value c"}
	// Value of pairs[3] is {"key b" "VALUE B"}
}

func ExampleMap() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key a", "value a")
	ctx = context.WithValue(ctx, "key b", "value b")
	ctx = context.WithValue(ctx, "key c", "value c")
	ctx = context.WithValue(ctx, "key b", "VALUE B")

	mapping := contents.Map(ctx)
	var lines []string

	for key, value := range mapping {
		lines = append(lines, fmt.Sprintf("Value of mapping[%q] is %q", key, value))
	}

	// This is only done because key order when ranging over a map is random
	sort.Strings(lines)

	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// Value of mapping["key a"] is "value a"
	// Value of mapping["key b"] is "VALUE B"
	// Value of mapping["key c"] is "value c"
}

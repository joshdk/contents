// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package contents

import (
	"context"
)

type Pair struct {
	Key   interface{}
	Value interface{}
}

// Keys will return every key contained withing the context, in the order in
// which they were originally added. Returned keys may be duplicates, but only
// because duplicates keys were added to the given context.
func Keys(ctx context.Context) []interface{} {
	var keys []interface{}

	// Do we have a parent context?
	parent := Unwrap(ctx)
	if parent != nil {
		// Extract keys from parent first
		keys = Keys(parent)
	}

	// Do we have a key?
	if key, found := Key(ctx); found {
		keys = append(keys, key)
	}

	return keys
}

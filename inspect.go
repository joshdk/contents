// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

// Package contents provides functions for extracting and walking the internals
// of a context.Context. This package imports both reflect and unsafe so use
// with caution.
//
// Internally, a context can be thought of as a singly-linked list, with a
// wrapper context holding a reference to the context that it wrapped. At
// a given level, an additional key:value pair may be attached.
//
//    context.Background() base context
//       ↑
//    context.WithValue("key-1", "val-1") adds the key "key-1"
//       ↑
//    context.WithCancel(...)
//       ↑
//    context.WithValue("key-2", "val-2") adds the key "key-2"
//       ↑
//    and so on
//
// The functions contained within this package care mostly about "Can this
// context level be unwrapped?" and "Does this context level contain a key?"
//
// This package also contains some functions that compose these two basic
// operations into helpers.
package contents

import (
	"context"
	"reflect"
	"unsafe"
)

// Unwrap takes a context and returns the wrapped context if it exists, and
// nil if it does not. A passed nil context will return nil.
//
// Contexts created with the "context.With___()" family of functions can be
// unwrapped, as they are derived from a "parent" context.
//
// Contexts created with "context.Background()" and "context.TODO()" can not be
// unwrapped, as they are not derived from a "parent" context.
func Unwrap(ctx context.Context) context.Context {

	// Guard against nil contexts
	if ctx == nil {
		return nil
	}

	contextVal := reflect.ValueOf(ctx).Elem()

	// Guard against types with no fields (such as context.Background)
	if contextVal.Kind() != reflect.Struct {
		return nil
	}

	// Obtain the struct field "Context"
	contextField := contextVal.FieldByName("Context")

	// Check to see if our field is actually a context
	if parentContext, ok := contextField.Interface().(context.Context); ok {
		return parentContext
	}

	return nil
}

// Key takes a context and returns the associated key and if a key exists.
// A passed nil context will return nil and false.
//
// Contexts created with "context.WithValue()" will have keys, but contexts
// created via other methods will not.
func Key(ctx context.Context) (interface{}, bool) {

	// Guard against nil contexts
	if ctx == nil {
		return nil, false
	}

	contextVal := reflect.ValueOf(ctx).Elem()

	// Guard against types with no fields (such as context.Background)
	if contextVal.Kind() != reflect.Struct {
		return nil, false
	}

	// Obtain the struct field "key"
	valueKey := contextVal.FieldByName("key")
	if valueKey.Kind() == reflect.Invalid {
		return nil, false
	}

	// Obtain a reference to the "key" field so that we can actually read its internal value
	refKey := reflect.NewAt(valueKey.Type(), unsafe.Pointer(valueKey.UnsafeAddr())).Elem()

	// Extract internal interface{} value
	return refKey.Interface(), true
}

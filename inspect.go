// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package contents

import (
	"context"
	"reflect"
	"unsafe"
)

func Unwrap(ctx context.Context) context.Context {

	if ctx == nil {
		return nil
	}

	contextVal := reflect.ValueOf(ctx).Elem()

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

func Key(ctx context.Context) (interface{}, bool) {

	// Guard against nil contexts
	if ctx == nil {
		return nil, false
	}

	contextVal := reflect.ValueOf(ctx).Elem()

	// Guard against types with no fields (such as context.Background)
	if contextVal.Kind().String() != "struct" {
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

[![License](https://img.shields.io/github/license/joshdk/contents.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/joshdk/contents?status.svg)](https://godoc.org/github.com/joshdk/contents)
[![Go Report Card](https://goreportcard.com/badge/github.com/joshdk/contents)](https://goreportcard.com/report/github.com/joshdk/contents)
[![CircleCI](https://circleci.com/gh/joshdk/contents.svg?&style=shield)](https://circleci.com/gh/joshdk/contents/tree/master)
[![CodeCov](https://codecov.io/gh/joshdk/contents/branch/master/graph/badge.svg)](https://codecov.io/gh/joshdk/contents)

# Contents

🔍 Inspect private context.Context internals

## Installing

You can fetch this library by running the following

    go get -u github.com/joshdk/contents

## Usage

```go
import (
	"context"
	"fmt"
	"github.com/joshdk/contents"
)

// Build a context
ctx := context.Background()
ctx = context.WithValue(ctx, "key-1", "val-1")
ctx = context.WithValue(ctx, "key-2", "val-2")
ctx = context.WithValue(ctx, "key-3", "val-3")

// Extract list of all keys
keys := contents.Keys(ctx)

for _, key := range keys {
	fmt.Printf("Context contains %q → %q\n", key, ctx.Value(key))
	// Context contains "key-1" → "val-1"
	// Context contains "key-2" → "val-2"
	// Context contains "key-3" → "val-3"
}
```

## License

This library is distributed under the [MIT License](https://opensource.org/licenses/MIT), see LICENSE.txt for more information.

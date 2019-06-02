package main

import (
	"context"
	"fmt"
)

func main() {
	type favContextKey string

	f := func(ctx context.Context, key favContextKey) {
		if value := ctx.Value(key); value != nil {
			fmt.Println("found value: ", value)
		} else {
			fmt.Println("key not found: ", key)
		}
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}

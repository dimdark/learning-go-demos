package main

import (
	"context"
	"fmt"
	"time"
)

func withTimeout() {
	// 50s
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func withTimeout2() {
	// 50ms
	ctx2, cancel2 := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel2()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx2.Done():
		fmt.Println(ctx2.Err())
	}
}

func main() {
	withTimeout()
	withTimeout2()
}

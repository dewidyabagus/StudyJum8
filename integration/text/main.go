package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type keyString string

const traceKey = keyString("trace-id")

func main() {
	// Parent -> Child

	// Empty context
	// Background()
	// Todo()

	// Dibagian turunannya ada
	// Value
	// Timeout / Deadline

	ctx := context.WithValue(context.Background(), traceKey, uuid.New().String()) // Setting value
	func1(ctx)
}

func func1(ctx context.Context) {
	traceID := ctx.Value(traceKey)
	fmt.Println("Run func1:", traceID)

	func2(ctx)
}

func func2(ctx context.Context) {
	traceID := ctx.Value(traceKey)
	fmt.Println("Run func2:", traceID)
}

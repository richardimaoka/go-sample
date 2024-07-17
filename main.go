// Option Pattern についてのサンプルです。
//
// #REFERENCES
//   - https://dev.to/c4r4x35/options-pattern-in-golang-10ph
package main

import (
	"context"
	"fmt"
)

type TraceID string
type AnotherID string

const ZeroTraceID = ""

type traceIDKey struct{}
type anotherIDKey struct{}

func SetTraceID(ctx context.Context, tid TraceID) context.Context {
	return context.WithValue(ctx, traceIDKey{}, tid)
}
func GetTraceID(ctx context.Context) TraceID {
	if v, ok := ctx.Value(traceIDKey{}).(TraceID); ok {
		return v
	}
	return ZeroTraceID
}

func SetAnotherID(ctx context.Context, aid AnotherID) context.Context {
	return context.WithValue(ctx, anotherIDKey{}, aid)
}
func GetAnotherID(ctx context.Context) AnotherID {
	if v, ok := ctx.Value(anotherIDKey{}).(AnotherID); ok {
		return v
	}
	return ZeroTraceID
}

func main() {
	ctx := context.Background()
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
	ctx = SetTraceID(ctx, "test-id")
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))

	fmt.Printf("another id = %q\n", GetAnotherID(ctx))
	ctx = SetAnotherID(ctx, "tessssst-id")
	fmt.Printf("another id = %q\n", GetAnotherID(ctx))
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
}

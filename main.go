// Option Pattern についてのサンプルです。
//
// #REFERENCES
//   - https://dev.to/c4r4x35/options-pattern-in-golang-10ph
package main

import (
	"context"
	"fmt"
)

func child(ctx context.Context) {
	// 関数の実ロジックに入る前にcontext.Contextの状態を検証する
	if err := ctx.Err(); err != nil {
		fmt.Println("キャンセルされている")
		return
	}
	fmt.Println("キャンセルされていない")
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx)
	cancel()
	child(ctx)
}

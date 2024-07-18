package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

// Common pattern : accept context.Context as the 1st argument.
func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	// 別ゴルーチンでHTTPサーバーを起動する
	eg.Go(func() error {
		// http.ErrServerClosed は
		// http.Server.Shutdown() が正常に終了したことを示すので異常ではない。
		fmt.Println("running HTTP server")
		if err := s.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	fmt.Println("waiting for a done notification")
	// チャネルからの通知（終了通知）を待機する
	<-ctx.Done()
	fmt.Println("received a done notification")
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	// Goメソッドで起動した別ゴルーチンの終了を待つ。
	return eg.Wait()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
		fmt.Println("canceled from a go routine")
	}()

	if err := run(ctx); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/richardimaoka/go-sandbox/config"
	"golang.org/x/sync/errgroup"
)

// Common pattern : accept context.Context as the 1st argument.
func run(ctx context.Context) error {
	//
	// Handling SIGNALs
	//
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	/*
	 * Load config for env name and port
	 */
	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("Start with: %v", url)

	s := &http.Server{
		// 引数で受け取ったnet.Listenerを利用するので、
		// Addrフィールドは指定しない
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// コマンドラインで実験するため
			time.Sleep(5 * time.Second)
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// ListenAndServeメソッドではなく、Serveメソッドに変更する
		fmt.Println("serving at ", l.Addr().String())
		if err := s.Serve(l); err != nil &&
			// http.ErrServerClosed は
			// http.Server.Shutdown() が正常に終了したことを示すので異常ではない
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
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

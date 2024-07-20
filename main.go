package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/richardimaoka/go-sandbox/config"
)

// Common pattern : accept context.Context as the 1st argument.
func run(ctx context.Context) error {
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

	mux := NewMux()
	s := NewServer(l, mux)
	return s.Run(ctx)
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

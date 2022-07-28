package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	gonnarain "github.com/100to-dev/go-daemon-example"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	cfg := &gonnarain.Config{}
	if err := cfg.Load(os.Args); err != nil {
		log.Fatalf("config load failed: %s", err)
	}

	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGINT, syscall.SIGTERM:
					log.Printf("got SIGINT/SIGTERM, exiting")
					cancel()
					os.Exit(1)
				case syscall.SIGHUP:
					log.Printf("got SIGHUP, reloading")
					cfg.Load(os.Args)
				}
			case <-ctx.Done():
				log.Printf("done")
				os.Exit(1)
			}
		}
	}()

	prefix := fmt.Sprintf("[gonnaraind] %d ", os.Getpid())
	logger := log.New(os.Stdout, prefix, log.Flags())

	if err := gonnarain.Run(ctx, cfg, logger); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}

package main

import (
	"context"
	"godemo/app"
	"godemo/runner"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	return runner.Run(ctx, &app.App{})
}

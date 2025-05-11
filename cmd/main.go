package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"webhook/internal/app"
	"webhook/internal/pkg/closer"
)

const bannerTemplate = `
                __    __                __  
 _      _____  / /_  / /_  ____  ____  / /__
| | /| / / _ \/ __ \/ __ \/ __ \/ __ \/ //_/
| |/ |/ /  __/ /_/ / / / / /_/ / /_/ / ,<   
|__/|__/\___/_.___/_/ /_/\____/\____/_/|_|  							   

ver: %s

`

var version = "local"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

		<-c
		cancel()
		slog.Info("application shutting down...")

		if err := closer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	application := app.NewApp()

	// nolint:forbidigo
	fmt.Printf(bannerTemplate, version)

	if err := application.Run(ctx); err != nil {
		log.Fatalf("fault run application err: %s", err.Error())
	}
}

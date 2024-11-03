// main.go
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/zuekveer/BWA/internal/application"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	if err := application.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

package application

import (
	"context"
	"fmt"
	"time"

	"github.com/zuekveer/BWA/internal/config"
	handlers "github.com/zuekveer/BWA/internal/controllers/http"
	"github.com/zuekveer/BWA/internal/databases"
	"github.com/zuekveer/BWA/internal/logger"
	"github.com/zuekveer/BWA/internal/transport/http"
)

const serverShutdownTimeout = 1 * time.Minute

func Run(ctx context.Context) error {
	cfg, err := config.Parse()
	if err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	log, err := logger.New()
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}
	_, close, err := databases.NewDB(&cfg.DB, log)
	if err != nil {
		return fmt.Errorf("failed to create databases: %w", err)
	}
	defer func() {
		err = close()
		if err != nil {
			log.Errorf("failed to close databases: %s", err)
		}
	}()

	hs := handlers.NewHandlers()
	stopHTTPServer, err := http.ServeHTTP(&cfg.HTTP, hs)
	if err != nil {
		return fmt.Errorf("failed to start HTTP server: %w", err)
	}
	// errors on server shutdown are not important
	// nolint:errcheck
	defer stopHTTPServer(func() context.Context {
		// context leaks on server shutdown are not important
		// nolint:govet
		ctx, _ := context.WithTimeout(context.Background(), serverShutdownTimeout)
		return ctx
	}())
	log.Infof("app started on port: %d", cfg.HTTP.Port)
	<-ctx.Done()
	return nil
}

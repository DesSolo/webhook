package app

import (
	"context"
	"fmt"
	"log/slog"

	"webhook/internal/pkg/closer"
)

// App application
type App struct{}

// New constructor
func New() *App {
	return &App{}
}

// Run application
func (a *App) Run(_ context.Context) error {
	// TODO: use context
	cfg, err := loadConfig()
	if err != nil {
		return fmt.Errorf("fault load config err: %w", err)
	}

	if err := configureLogger(cfg); err != nil {
		return fmt.Errorf("fault configure logger err: %w", err)
	}

	ps, err := loadPubSub(cfg)
	if err != nil {
		return fmt.Errorf("fault load pubsub err: %w", err)
	}

	closer.Add(ps.Close)

	st, err := loadResponseStorage(cfg)
	if err != nil {
		return fmt.Errorf("fault load storageerr: %w", err)
	}

	closer.Add(st.Close)

	ws, err := loadWebhookService(ps, st)
	if err != nil {
		return fmt.Errorf("fault load webhook serviceerr: %w", err)
	}

	srv, err := loadServer(cfg, ps, ws)
	if err != nil {
		return fmt.Errorf("fault load servererr: %w", err)
	}

	slog.Info("server running", "address", cfg.BindAddress)

	if err := srv.Run(cfg.BindAddress); err != nil {
		return fmt.Errorf("fault run server err: %w", err)
	}

	return nil
}

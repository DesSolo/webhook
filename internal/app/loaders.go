package app

import (
	"fmt"
	"log/slog"
	"os"

	goredis "github.com/redis/go-redis/v9"

	"webhook/internal/config"
	"webhook/internal/pkg/logger"
	"webhook/internal/pkg/pubsub"
	"webhook/internal/pkg/pubsub/channel"
	predis "webhook/internal/pkg/pubsub/redis"
	"webhook/internal/pkg/server"
	"webhook/internal/pkg/service"
	"webhook/internal/pkg/storage"
	"webhook/internal/pkg/storage/memory"
	sredis "webhook/internal/pkg/storage/redis"
)

// loadConfig load configuration
func loadConfig() (*config.Config, error) {
	path := os.Getenv("CONFIG_FILE_PATH")
	if path == "" {
		path = "/etc/webhook/config.yml"
	}

	return config.FromFile(path)
}

// configureLogger configure logger
func configureLogger(cfg *config.Config) error {
	o := slog.HandlerOptions{
		AddSource: cfg.Logging.Option.AddSource,
		Level:     slog.Level(cfg.Logging.Option.Level),
	}

	var h slog.Handler

	switch cfg.Logging.Handler {
	case "text":
		h = slog.NewTextHandler(os.Stdout, &o)
	case "json":
		h = slog.NewJSONHandler(os.Stdout, &o)
	default:
		return fmt.Errorf("handler: %q not supported", cfg.Logging.Handler)
	}

	args := make([]any, 0, len(cfg.Logging.Args))

	for k, v := range cfg.Logging.Args {
		args = append(args, k, v)
	}

	h = logger.NewLogContextHandler(h)

	slog.SetDefault(
		slog.New(h).With(args...),
	)

	return nil
}

// loadPubSub load pubsub
func loadPubSub(cfg *config.Config) (pubsub.PubSub, error) {
	switch cfg.PubSub.Kind {
	case "channel":
		return channel.New(), nil
	case "redis":
		c := goredis.NewClient(&goredis.Options{
			Addr: cfg.PubSub.Redis.Addr,
			DB:   cfg.PubSub.Redis.DB,
		})
		return predis.New(c), nil
	default:
		return nil, fmt.Errorf("pubsub: %q not supported", cfg.PubSub.Kind)
	}
}

// loadResponseStorage load response storage
func loadResponseStorage(cfg *config.Config) (storage.ResponseStorage, error) {
	switch cfg.Storage.Kind {
	case "memory":
		return memory.New(), nil
	case "redis":
		c := goredis.NewClient(&goredis.Options{
			Addr: cfg.Storage.Redis.Addr,
			DB:   cfg.Storage.Redis.DB,
		})
		return sredis.New(c), nil
	default:
		return nil, fmt.Errorf("storage: %q not supported", cfg.Storage.Kind)
	}
}

// loadWebhookService load webhook service
func loadWebhookService(ps pubsub.PubSub, st storage.ResponseStorage) (*service.Webhook, error) {
	return service.NewWebhook(ps, st), nil
}

// loadServer load http server
func loadServer(cfg *config.Config, ps pubsub.PubSub, ws *service.Webhook) (*server.Server, error) {
	srv := server.NewServer(server.Options{
		ServeStatic: cfg.Server.ServeStatic,
		StaticPath:  cfg.Server.StaticPath,
	})
	srv.LoadRoutes(ps, ws)

	return srv, nil
}

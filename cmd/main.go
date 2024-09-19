package main

import (
	"fmt"
	"log/slog"
	"os"
	"webhook/config"
	"webhook/internal/logger"
	"webhook/internal/pubsub"
	"webhook/internal/pubsub/channel"
	predis "webhook/internal/pubsub/redis"
	"webhook/internal/server"
	"webhook/internal/service"
	"webhook/internal/storage"
	"webhook/internal/storage/memory"
	sredis "webhook/internal/storage/redis"

	goredis "github.com/redis/go-redis/v9"
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
func loadWebhookService(_ *config.Config, ps pubsub.PubSub, st storage.ResponseStorage) (*service.Webhook, error) {
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

// fatal force exit with error
func fatal(message string, err error) {
	slog.Error(message, "err", err)
	os.Exit(1)
}

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
	cfg, err := loadConfig()
	if err != nil {
		fatal("fault load config", err)
	}

	if err := configureLogger(cfg); err != nil {
		fatal("fault configure logger", err)
	}

	ps, err := loadPubSub(cfg)
	if err != nil {
		fatal("fault load pubsub", err)
	}

	st, err := loadResponseStorage(cfg)
	if err != nil {
		fatal("fault load storage", err)
	}

	ws, err := loadWebhookService(cfg, ps, st)
	if err != nil {
		fatal("fault load webhook service", err)
	}

	srv, err := loadServer(cfg, ps, ws)
	if err != nil {
		fatal("fault load server", err)
	}

	fmt.Printf(bannerTemplate, version)

	slog.Info("server running", "address", cfg.BindAddress)

	if err := srv.Run(cfg.BindAddress); err != nil {
		fatal("fault run server", err)
	}
}

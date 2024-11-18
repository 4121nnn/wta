package logger

import (
	"log"
	"log/slog"
	"os"
	"sync"
	"wta/internal/logger/handlers/slogpretty"
)

const (
	LOCAL = "local"
	DEV   = "dev"
	PROD  = "prod"
)

var (
	logger *slog.Logger
	once   sync.Once
)

func Setup(env string) {
	once.Do(func() {
		switch env {
		case LOCAL:
			logger = setupPrettySlog()
		case DEV:
			logger = slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)
		case PROD:
			logger = slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
			)
		}
	})
}

func Get() *slog.Logger {
	if logger == nil {
		log.Fatal("logger not initialized")
	}
	return logger
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

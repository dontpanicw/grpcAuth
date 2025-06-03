package main

import (
	"fmt"
	"log/slog"
	"os"
	"sso/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//TODO: инициализация конфига
	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("starting application",
		//slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
		//slog.Int("port", cfg.GRPC.Port),
	)

	//log.Debug("debug message")
	//
	//log.Error("error message")
	//
	//log.Warn("warning message")

	//TODO: логирование

	//TODO: инициализация приложения

	//TODO: запустить grpc сервер приложения
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

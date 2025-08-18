package main

import (
	"log/slog"
	"net/http"
	"os"

	"url-shortener/src/config"
	"url-shortener/src/internal/adapter"
	"url-shortener/src/internal/data/repository"
	"url-shortener/src/internal/domain/service"
	"url-shortener/src/internal/handler"
	"url-shortener/src/internal/web"
	"url-shortener/src/pkg/shortener"

	"github.com/go-chi/chi/v5"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg := config.MustLoad()

	log := SetupLogger(cfg.Env)
	log.Info("Start", slog.String("env", cfg.Env))
	log.Debug("debag enable")

	repo, err := repository.NewPostgresRepo(cfg.DatabaseURL, log)
	if err != nil {
		log.Error("cannot connect to db", slog.Any("err", err))
		return
	}
	adaRepo := &adapter.PostgresAdapter{Repo: repo}

	svc := service.NewURLService(adaRepo)
	h := handler.New(svc, shortener.GenerateID)

	webHandler := web.NewWebHandler(svc, shortener.GenerateID)

	r := chi.NewRouter()
	r.Post("/shorten", h.ShortenURL)
	r.Get("/{id}", h.Redirect)

	r.Get("/", webHandler.Home)
	r.Post("/shorten-web", webHandler.ShortenWeb)

	http.ListenAndServe(":8080", r)
}

func SetupLogger(env string) *slog.Logger {
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

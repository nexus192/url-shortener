package main

import (
	"log"
	_ "log/slog"
	"net/http"
	"time"

	"url-shortener/internal/data/repository"
	"url-shortener/internal/domain/service"
	"url-shortener/internal/handler"
	"url-shortener/internal/web"
	"url-shortener/pkg/shortener"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	repo := repository.NewInMemoryRepo()
	svc := service.NewURLService(repo)
	h := handler.New(svc, shortener.GenerateID)

	webHandler := web.NewWebHandler(svc, shortener.GenerateID)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Post("/shorten", h.ShortenURL)
	r.Get("/{id}", h.Redirect)

	r.Get("/", webHandler.Home)
	r.Post("/shorten-web", webHandler.ShortenWeb)

	log.Println("Server is running on :8080")
	http.ListenAndServe(":8080", r)
}

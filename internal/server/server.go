package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/vitorvargasdev/animefluency-go-api/internal/router"
)

func StartServer(port int) error {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error loading .env file: %w", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.AddRoutes(r)

	portString := fmt.Sprintf(":%d", port)

	if err := http.ListenAndServe(portString, r); err != nil {
		return err
	}

	return nil
}

package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/vitorvargasdev/animefluency-go-api/internal/database"
	"github.com/vitorvargasdev/animefluency-go-api/internal/router"
	"github.com/vitorvargasdev/animefluency-go-api/internal/service"
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

	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := service.NewPool(db); err != nil {
		log.Fatal(err)
	}

	router.AddRoutes(r)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		return err
	}

	return nil
}

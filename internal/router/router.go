package router

import (
	"github.com/go-chi/chi/v5"
)

func AddRoutes(r *chi.Mux) {
	r.Get("/version", getVersion)
}

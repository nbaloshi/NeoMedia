package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func HealthRoute(r *chi.Mux, db *pgx.Conn) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		// Simple health response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}
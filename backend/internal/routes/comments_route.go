package routes

import (
	"NeoMedia/internal/handlers/comments"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CommentsRoutes(r *chi.Mux, db *pgxpool.Pool) {
	r.Post("/comments", func(w http.ResponseWriter, r *http.Request) {
		comments.CommentsCreateHandler(w, r, db)
	})

	r.Get("/comments", func(w http.ResponseWriter, r *http.Request) {
		comments.CommentsHandler(w, r, db)
	})
}
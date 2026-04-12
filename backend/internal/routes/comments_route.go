package routes

import (
	"NeoMedia/internal/handlers/comments"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func CommentsRoutes(r *chi.Mux, db *pgx.Conn) {
	r.Post("/comments", func(w http.ResponseWriter, r *http.Request) {
		comments.CommentsCreateHandler(w, r, db)
	})
}
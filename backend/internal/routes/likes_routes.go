package routes

import (
	"NeoMedia/internal/handlers/likes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func LikesRoutes(r *chi.Mux, db *pgxpool.Pool) {
	r.Post("/likes-post", func(w http.ResponseWriter, r *http.Request) {
		likes.LikesPostsToggleHandler(w, r, db)
	})

	r.Get("/likes-post", func(w http.ResponseWriter, r *http.Request) {
		likes.LikesPostsHandler(w, r, db)
	})
}
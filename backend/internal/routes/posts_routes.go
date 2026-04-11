package routes

import (
	"NeoMedia/internal/handlers/posts"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func PostsRoutes(r *chi.Mux, db *pgx.Conn) {
	r.Post("/posts", func(w http.ResponseWriter, r *http.Request) {
		posts.PostsCreateHandler(w, r, db)
	})

	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		posts.PostsHandler(w, r, db)
	})
}
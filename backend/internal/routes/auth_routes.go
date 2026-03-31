package routes

import (
	"NeoMedia/internal/handlers/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func AuthRoutes(r *chi.Mux, db *pgx.Conn) {
	// Registration route
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		auth.RegisterationHandler(w, r, db)
	})

	// Login route
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		auth.LoginHandler(w, r, db)
	})

	// Logout route
	r.Post("/logout", func(w http.ResponseWriter, r *http.Request) {
		auth.LogoutHandler(w, r, db)
	})
	
	// User route
	r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
		auth.MeHandler(w, r, db)
	})
}
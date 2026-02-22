package main

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/auth"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Root context that cancels on SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Connect to databse
	connStr := "postgres://neouser:84560505@localhost:5432/neomedia?sslmode=disable"
	db, err := sql.GetDB(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP server with /health route
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		// Simple health response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

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

	// Goroutine server
	server := &http.Server{
		Addr:	":8080",
		Handler: r,
	}
	go func()  {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
	fmt.Println("Server is running on http://localhost:8080")

	// Blockade until shutdown signal
	<-ctx.Done()
	fmt.Println("\nShutdown signal received")

	//Gracefully close DB connection with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sql.CloseDB(shutdownCtx, db)
	fmt.Println("Program exited cleanly")
}
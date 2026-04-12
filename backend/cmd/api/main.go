package main

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/routes"
	"NeoMedia/internal/sessions"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	// Root context that cancels on SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Connect to databse
	connStr := "postgres://neouser:84560505@localhost:5432/neomedia?sslmode=disable"
	db, err := sql.GetDB(ctx, connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Setup HTTP server with /health route
	r := chi.NewRouter()

	// Global middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: 	[]string{"http://localhost:5173"},
		AllowedMethods: 	[]string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: 	[]string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: 	true,
	}))
	r.Use(sessions.Middleware(db))

	// Routes
	routes.HealthRoute(r, db)
	routes.AuthRoutes(r, db)
	routes.PostsRoutes(r, db)
	routes.CommentsRoutes(r, db)

	// Goroutine server
	server := &http.Server{
		Addr:	":8080",
		Handler: r,
	}
	go func()  {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
			stop()
		}
	}()
	fmt.Println("Server is running on http://localhost:8080")

	// Blockade until shutdown signal
	<-ctx.Done()
	fmt.Println("Shutdown signal received")

	//Gracefully close DB connection with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	sql.CloseDB(shutdownCtx, db)
	fmt.Println("Program exited cleanly")
}
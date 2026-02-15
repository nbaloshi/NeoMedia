package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func main() {
	// Root context that cancels on SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Connect to databse
	connStr := "postgres://neouser:84560505@localhost:5432/neomedia?sslmode=disable"
	db, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal("Unable to connect to databse:", err)
	}
	fmt.Println("Connected to database")

	// Setup HTTP server with /health route
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		// Simple health response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
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
	if err := db.Close(shutdownCtx); err != nil {
		log.Printf("Error closing DB connection: %v", err)
	} else {
		fmt.Println("Disconnected from database")
	}
	fmt.Println("Program exited cleanly")
}
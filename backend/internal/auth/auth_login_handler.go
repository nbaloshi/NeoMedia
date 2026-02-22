package auth

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/sessions"
	"NeoMedia/internal/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type LoginBody struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	var body LoginBody

	// decode from frontend and put it in the body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("Login handler - Error decoding request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// loop through database and match with the email and fetch email and hashed password
	userQuery := `SELECT id, password_hash FROM users WHERE email = $1`
	
	var userId int
	var hashedPassword string

	querErr := sql.QueryRow(r.Context(), conn, userQuery,
		body.Email,
	).Scan(&userId, &hashedPassword)
	if err == pgx.ErrNoRows {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	} else if querErr != nil {
        log.Printf("Login handler - Error fetching user: %v", querErr)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

	// compare the hashed password with password
	if !utils.ComparePasswordHash(body.Password, hashedPassword) {
		log.Printf("Login handler - hash compare false")
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
	}

	// delete previous session for the user
	if err := sessions.DeleteSessionByUserId(r, conn, userId); err != nil {
		http.Error(w, "Failed to clear old sessions", http.StatusInternalServerError)
		return
	}

	// create a session token and store with user id in the sessions table
	if err := sessions.CreateSession(r, conn, userId); err != nil {
		http.Error(w, "Failed to create sessions", http.StatusInternalServerError)
		return
	}

	// respond
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
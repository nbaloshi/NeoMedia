package auth

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/sessions"
	"NeoMedia/internal/utils"
	"encoding/json"
	"fmt"
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
		log.Printf("LoginHandler - Error decoding request body: %v", err)
		utils.RespondError(w, http.StatusBadRequest, "Bad request")
		return
	}

	// loop through database and match with the email and fetch email and hashed password
	userQuery := `SELECT id, password_hash FROM users WHERE email = $1`
	
	var userId string
	var hashedPassword string

	querErr := sql.QueryRow(r.Context(), conn, userQuery,
		body.Email,
	).Scan(&userId, &hashedPassword)
	if querErr == pgx.ErrNoRows {
		utils.RespondError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	} else if querErr != nil {
        log.Printf("LoginHandler - DB query error: %v", querErr)
        utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
        return
    }

	// compare the hashed password with password
	if !utils.ComparePasswordHash(body.Password, hashedPassword) {
		log.Printf("LoginHandler - Password mismatch for email: %s", body.Email)
        utils.RespondError(w, http.StatusUnauthorized, "Invalid email or password")
        return
	}

	// create session
	if err := sessions.StartSession(w, r, conn, userId); err != nil {
		log.Printf("LoginHandler - Session creation failed: %v", err)
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create session")
		return
	}
	fmt.Println("LoginHandler - Login successful")

	// respond
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Login successful"}`))
}
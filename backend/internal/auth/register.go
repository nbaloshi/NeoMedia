package auth

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

type RegisterBody struct {
	FirstName   string   	`json:"firstname"`
	LastName    string   	`json:"lastname"`
	Email       string   	`json:"email"`
	Gender      string   	`json:"gender"`
	DateOfBirth time.Time 	`json:"dateOfBirth"`
	UserName    string   	`json:"username"`
	Password    string   	`json:"password"`
	ProfilePic  string   	`json:"profilePic"`
}

func RegisterationHandler(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	var body RegisterBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("Registration handler - Error decoding request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	hashedPassword, hasherr := utils.HashPassword(body.Password)
	if hasherr != nil {
		log.Printf("Registration handler - Error hashing password: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO users (firstname, lastname, email, gender, date_of_birth, username, password_hash, profile_pic)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`

	execErr := sql.Exec(r.Context(), conn, query,
		body.FirstName,
		body.LastName,
		body.Email,
		body.Gender,
		body.DateOfBirth,
		body.UserName,
		hashedPassword,
		body.ProfilePic,
	)

	if execErr != nil {
        log.Printf("Error inserting user: %v", execErr)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}
package auth

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RegisterBody struct {
	FirstName   string   	`json:"firstname"`
	LastName    string   	`json:"lastname"`
	Email       string   	`json:"email"`
	Gender      string   	`json:"gender"`
	DateOfBirth string 		`json:"dateOfBirth"`
	UserName    string   	`json:"username"`
	Password    string   	`json:"password"`
	ProfilePic  string   	`json:"profilePic"`
}

func RegisterationHandler(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
	var body RegisterBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("RegistrationHandler - Error decoding body: %v", err)
		utils.RespondError(w, http.StatusBadRequest, "Bad request")
		return
	}

	hashedPassword, hashErr := utils.HashPassword(body.Password)
	if hashErr != nil {
		log.Printf("RegistrationHandler - Error hashing password: %v", hashErr)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	query := `
		INSERT INTO users (firstname, lastname, email, gender, date_of_birth, username, password_hash, profile_pic)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`

	execErr := sql.Exec(r.Context(), pool, query,
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
        log.Printf("RegistrationHandler - Error inserting user: %v", execErr)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
		return
    }

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"User registered successfully"}`))
}
package auth

import (
	"NeoMedia/internal/sessions"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type TokenBody struct {
	TokenSession string `json:"sessionToken"`
}

func LogoutHandler(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	var body TokenBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("Logout handler - Error decoding request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	token := body.TokenSession

	if err := sessions.DeleteSessionByToken(r, conn, token); err != nil {
		log.Printf("Logout handler - Error deleting session: %v", err)
		http.Error(w, "Failed to clear session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
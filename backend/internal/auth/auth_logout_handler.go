package auth

import (
	"NeoMedia/internal/sessions"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {

	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "No session cookie", http.StatusUnauthorized)
		return
	}

	token := cookie.Value

	if err := sessions.EndSession(w, r, conn, token); err != nil {
		log.Printf("Logout handler - Error deleting session: %v", err)
		http.Error(w, "Failed to clear session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
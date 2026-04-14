package auth

import (
	"NeoMedia/internal/sessions"
	"NeoMedia/internal/utils"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
	token, ok := sessions.GetUserTokenFromContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := sessions.EndSession(w, r, pool, token); err != nil {
		log.Printf("LogoutHandler - Error ending session: %v", err)
		utils.RespondError(w, http.StatusInternalServerError, "Failed to clear session")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Logout successful"}`))
}
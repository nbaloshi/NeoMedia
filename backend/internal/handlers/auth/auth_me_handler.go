package auth

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/sessions"
	"NeoMedia/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MeHandler(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
	userId, ok := sessions.GetUserIDFromContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "Unauthorized")
		fmt.Println("issue here")
		return
	}

	var email string
	var username string
	query := `SELECT email, username FROM users WHERE id = $1`
	err := sql.QueryRow(r.Context(), pool, query, userId).Scan(&email, &username)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to fetch user")
		return
	}

	resp := map[string]interface{}{
		"id":	 userId,
		"email": email,
		"username": username,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
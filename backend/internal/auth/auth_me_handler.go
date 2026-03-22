package auth

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/sessions"
	"NeoMedia/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func MeHandler(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	userId, ok := sessions.GetUserIDFromContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "Unauthorized")
		fmt.Println("issue here")
		return
	}

	var email string
	query := `SELECT email FROM users WHERE id = $1`
	err := sql.QueryRow(r.Context(), conn, query, userId).Scan(&email)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to fetch user")
		return
	}

	resp := map[string]interface{}{
		"id":	 userId,
		"email": email,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
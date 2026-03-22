package sessions

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

func CreateSession(r *http.Request, conn *pgx.Conn, userId string) (string, error) {
	token, err := utils.GenerateSessionToken(16)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	query := `
		INSERT INTO sessions (user_id, session_token, expires_at)
		VALUES ($1, $2, $3)
	`

	execErr := sql.Exec(r.Context(), conn, query,
		userId,
		token,
		time.Now().Add(30 * time.Minute),
	)
	if execErr != nil {
        return "", fmt.Errorf("failed to execute query: %w", execErr)
    }

	return token, nil
}

func FetchSession(r *http.Request, conn *pgx.Conn, token string) (string, error) {
	var userId string
	var expiresAt time.Time

	query := `SELECT user_id, expires_at FROM sessions WHERE session_token = $1`
	queryErr := sql.QueryRow(r.Context(), conn, query, token).Scan(&userId, &expiresAt)
	if queryErr != nil {
		return "", fmt.Errorf("failed to fetch session: %w", queryErr)
	}

	if utils.HasExpired(expiresAt) {
		return "", fmt.Errorf("Session expired")
	}

	return userId, nil
}

func DeleteSessionByUserId(r *http.Request, conn *pgx.Conn, userId string) error {
	query := `DELETE FROM sessions WHERE user_id = $1`
	queryErr := sql.Exec(r.Context(), conn, query, userId)
	if queryErr != nil {
		return fmt.Errorf("failed to delete session(s): %w", queryErr)
	}
	return nil
}

func DeleteSessionByToken(r *http.Request, conn *pgx.Conn, token string) error {
	query := `DELETE FROM sessions WHERE session_token = $1`
	queryErr := sql.Exec(r.Context(), conn, query, token)
	if queryErr != nil {
		return fmt.Errorf("failed to delete session(s): %w", queryErr)
	}
	return nil
}
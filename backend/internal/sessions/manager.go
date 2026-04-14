package sessions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StartSession(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool, userId string) error {
	// Clean up old sessions
	if err := DeleteSessionByUserId(r, pool, userId); err != nil {
		return err
	}
	fmt.Println("StartSession - old session deleted")

	// Create new session in DB
	token, err := CreateSession(r, pool, userId)
	if err != nil {
		return err
	}

	// Set cookie in browser directly
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  time.Now().Add(30 * time.Minute),
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true, // only in production with HTTPS
		Path:     "/",
	}
	http.SetCookie(w, cookie)

	return nil
}

func EndSession(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool, token string) error {
	if err := DeleteSessionByToken(r, pool, token); err != nil {
		return err
	}

	// Clear cookie
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)

	return nil
}

// GetUserFromSession validates a session cookie and returns user ID
func GetUserFromSession(r *http.Request, pool *pgxpool.Pool) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return "", err
	}

	userId, err := FetchSession(r, pool, cookie.Value)
	if err != nil {
		return "", err
	}

	return userId, nil
}
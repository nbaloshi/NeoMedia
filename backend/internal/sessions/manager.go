package sessions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

func StartSession(w http.ResponseWriter, r *http.Request, conn *pgx.Conn, userId string) error {
	// Clean up old sessions
	if err := DeleteSessionByUserId(r, conn, userId); err != nil {
		return err
	}
	fmt.Println("StartSession - old session deleted")

	// Create new session in DB
	token, err := CreateSession(r, conn, userId)
	if err != nil {
		return err
	}

	// Set cookie in browser directly
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  time.Now().Add(30 * time.Minute),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   true, // only in production with HTTPS
		Path:     "/",
	}
	http.SetCookie(w, cookie)

	return nil
}

func EndSession(w http.ResponseWriter, r *http.Request, conn *pgx.Conn, token string) error {
	if err := DeleteSessionByToken(r, conn, token); err != nil {
		return err
	}

	// Clear cookie
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)

	return nil
}

// GetUserFromSession validates a session cookie and returns user ID
func GetUserFromSession(r *http.Request, conn *pgx.Conn) (int, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, err
	}

	userId, err := FetchSession(r, conn, cookie.Value)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
package sessions

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/utils"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ContextKey string

const userIdKey ContextKey = "userId"
const usernameKey ContextKey = "username"
const userTokenKey ContextKey = "token_session"

func Middleware(pool *pgxpool.Pool) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		handler := func(w http.ResponseWriter, r *http.Request) {
			// white list check
			if r.URL.Path == "/health" || r.URL.Path == "/register" || r.URL.Path == "/login" {
				next.ServeHTTP(w, r)
				return
			}

			// Step 1: Read cookie
			cookie, err := r.Cookie("session_token")
			if err != nil {
				http.Error(w, "Unauthorized: no session cookie", http.StatusUnauthorized)
 				return
			}
			token := cookie.Value

			// Step 2: Fetch session from DB
			var userId string
			var expiresAt time.Time
			fetchQuery := `SELECT user_id, expires_at FROM sessions WHERE session_token = $1 AND is_active = true`
			fetchQueryErr := sql.QueryRow(r.Context(), pool, fetchQuery, token).Scan(&userId, &expiresAt)
			if fetchQueryErr != nil {
				http.Error(w, "Unauthorized: invalid session", http.StatusUnauthorized)
				return
			}

			// Step 3: Expiry check
			if utils.HasExpired(expiresAt) {
				http.Error(w, "Unauthorized: session expired", http.StatusUnauthorized)
				return
			}

			// Step 4: Sliding expiration (optional)
			updateQuery := `UPDATE sessions SET last_accessed = $1, expires_at = $2 WHERE session_token = $3`
			updateQueryErr := sql.Exec(r.Context(), pool, updateQuery,
				time.Now(),
				time.Now().Add(30 * time.Minute),
				token,
			)
			if updateQueryErr != nil {
				log.Printf("AuthMiddleware - failed to update session: %v", updateQueryErr)
			}

			// Step 5: Refresh cookie expiry in browser
			refreshCookie := &http.Cookie{
				Name: "session_token",
				Value: token,
				Expires: time.Now().Add(30 * time.Minute),
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
				Secure: true,
				Path: "/",
			}
			http.SetCookie(w, refreshCookie)

			// Step 6: Attach userId and token to context
			ctx := context.WithValue(r.Context(), userIdKey, userId)
			ctx = context.WithValue(ctx, userTokenKey, token)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(handler)
	}
}

func GetUserIDFromContext(r *http.Request) (string, bool) {
	userId, ok := r.Context().Value(userIdKey).(string)
	return userId, ok
}

func GetUsernameFromContext(r *http.Request) (string, bool) {
	username, ok := r.Context().Value(usernameKey).(string)
	return username, ok
}

func GetUserTokenFromContext(r *http.Request) (string, bool) {
	token, ok := r.Context().Value(userTokenKey).(string)
	return token, ok
}
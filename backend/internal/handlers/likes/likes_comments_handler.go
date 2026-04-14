package likes

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/sessions"
	"NeoMedia/internal/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func LikesCommentsHandler(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
	userId, ok := sessions.GetUserIDFromContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "Unauthorized")
        return
	}

	commentId := r.URL.Query().Get("comment_id")
	if commentId == "" {
        utils.RespondError(w, http.StatusBadRequest, "Missing comment_id")
        return
    }

	var count int
	queryCount := `SELECT COUNT(*) FROM comments_likes WHERE comment_id=$1`
	err := sql.QueryRow(r.Context(), pool, queryCount, commentId).Scan(&count)
	if err != nil {
		log.Printf("LikesCommentsHandler - Error counting like: %v", err)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
        return
	}

	var likedByMe bool
	queryLiked := `SELECT EXISTS(SELECT 1 FROM comments_likes WHERE comment_id=$1 AND user_id=$2)`
	errLike := sql.QueryRow(r.Context(), pool, queryLiked, commentId, userId).Scan(&likedByMe)
	if errLike != nil {
		log.Printf("LikesCommentsHandler - Error checking like: %v", err)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
        return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"count":	 count,
		"likedByMe": likedByMe,
	})
}
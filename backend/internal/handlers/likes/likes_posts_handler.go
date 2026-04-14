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

func LikesPostsHandler(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
	userId, ok := sessions.GetUserIDFromContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "Unauthorized")
        return
	}

	postId := r.URL.Query().Get("post_id")
	if postId == "" {
        utils.RespondError(w, http.StatusBadRequest, "Missing post_id")
        return
    }

	var count int
	queryCount := `SELECT COUNT(*) FROM post_likes WHERE post_id=$1`
	err := sql.QueryRow(r.Context(), pool, queryCount, postId).Scan(&count)
	if err != nil {
		log.Printf("LikesPostsHandler - Error counting like: %v", err)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
        return
	}

	var likedByMe bool
	queryLiked := `SELECT EXISTS(SELECT 1 FROM post_likes WHERE post_id=$1 AND user_id=$2)`
	errLike := sql.QueryRow(r.Context(), pool, queryLiked, postId, userId).Scan(&likedByMe)
	if errLike != nil {
		log.Printf("LikesPostsHandler - Error checking like: %v", err)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
        return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"count":	 count,
		"likedByMe": likedByMe,
	})
}
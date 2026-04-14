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

func LikesCommentsToggleHandler(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
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

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM comments_likes WHERE comment_id=$1 AND user_id=$2)`
	err := sql.QueryRow(r.Context(), pool, query, commentId, userId).Scan(&exists)
	if err != nil {
		log.Printf("LikesCommentsToggleHandler - Error checking exists: %v", err)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
        return
	}

	if exists {
		query := `DELETE FROM comments_likes WHERE comment_id=$1 AND user_id=$2`
		err := sql.Exec(r.Context(), pool, query, commentId, userId)
		if err != nil {
			log.Printf("LikesCommentsToggleHandler - Error deleting like: %v", err)
			utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
            return
		}
		json.NewEncoder(w).Encode(map[string]any{"liked": false})
	} else {
		query := `INSERT INTO comments_likes (comment_id, user_id) VALUES ($1, $2)`
		err := sql.Exec(r.Context(), pool, query, commentId, userId)
		if err != nil {
			log.Printf("LikescommentsToggleHandler - Error inserting like: %v", err)
			utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
            return
		}
		json.NewEncoder(w).Encode(map[string]any{"liked": true})
	}
		
}
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

func LikesPostsToggleHandler(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
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

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM post_likes WHERE post_id=$1 AND user_id=$2)`
	err := sql.QueryRow(r.Context(), pool, query, postId, userId).Scan(&exists)
	if err != nil {
		log.Printf("LikesPostsToggleHandler - Error checking exists: %v", err)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
        return
	}

	if exists {
		query := `DELETE FROM post_likes WHERE post_id=$1 AND user_id=$2`
		err := sql.Exec(r.Context(), pool, query, postId, userId)
		if err != nil {
			log.Printf("LikesPostsToggleHandler - Error deleting like: %v", err)
			utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
            return
		}
		json.NewEncoder(w).Encode(map[string]any{"liked": false})
	} else {
		query := `INSERT INTO post_likes (post_id, user_id) VALUES ($1, $2)`
		err := sql.Exec(r.Context(), pool, query, postId, userId)
		if err != nil {
			log.Printf("LikesPostsToggleHandler - Error inserting like: %v", err)
			utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
            return
		}
		json.NewEncoder(w).Encode(map[string]any{"liked": true})
	}
		
}
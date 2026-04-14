package comments

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Comment struct {
	ID 			string `json:"id"`
	Username	string 	`json:"username"`
	Content 	string `json:"content"`
	CreatedAt 	time.Time `json:"createdAt"`
}

func CommentsHandler(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
	postId := r.URL.Query().Get("post_id")
	if postId == "" {
		utils.RespondError(w, http.StatusBadRequest, "Missing post_id")
    	return
	}

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	limit := 10
	page := 0

	if p, err := strconv.Atoi(pageStr); err == nil {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil {
		limit = l
	}

	offset := page * limit

	query := `SELECT id, username, content, created_at
		FROM comments
		WHERE post_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, querryErr := sql.Query(r.Context(), pool, query, postId, limit, offset)
	if querryErr != nil {
		log.Printf("CommentsHandler - Error querrying comments: %v", querryErr)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var c Comment
		if err := rows.Scan(&c.ID, &c.Username, &c.Content, &c.CreatedAt); err != nil {
			log.Printf("CommentsHandler - Error scanning row: %v", err)
            continue
		}
		comments = append(comments, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
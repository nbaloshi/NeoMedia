package posts

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

type Post struct {
	ID		  	string 	`json:"id"`
	Username	string 	`json:"username"`
	Content	  	string 	`json:"content"`
	CreatedAt 	time.Time `json:"createdAt"`
}

func PostsHandler(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 0
	limit := 10
	if p, err := strconv.Atoi(pageStr); err == nil {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil {
		limit = l
	}

	offset := page * limit

	query := `SELECT id, username, content, created_at
			  FROM posts
			  ORDER BY created_at DESC
			  LIMIT $1 OFFSET $2
	`

	rows, querryErr := sql.Query(r.Context(), conn, query, limit, offset)
	if querryErr != nil {
		log.Printf("PostsHandler - Error querrying posts: %v", querryErr)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Username, &p.Content, &p.CreatedAt); err != nil {
			log.Printf("PostsHandler - Error scanning row: %v", err)
            continue
		}
		posts = append(posts, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
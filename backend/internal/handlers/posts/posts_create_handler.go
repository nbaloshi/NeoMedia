package posts

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/sessions"
	"NeoMedia/internal/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostCreateBody struct {
	Content string `json:"content"`
}

func PostsCreateHandler(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
	userId, ok := sessions.GetUserIDFromContext(r)
		if !ok {
    	utils.RespondError(w, http.StatusUnauthorized, "Unauthorized")
    	return
	}

	var body PostCreateBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("CreatePostsHandler - Error decoding body: %v", err)
		utils.RespondError(w, http.StatusBadRequest, "Bad request")
		return
	}

	var username string
	queryUsername := `SELECT username FROM users WHERE id = $1`
	queryErr := sql.QueryRow(r.Context(), pool, queryUsername, userId).Scan(&username)
	if queryErr != nil {
		log.Printf("CreatePostsHandler - Error querrying username: %v", queryErr)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	query := `INSERT INTO posts (username, content) VALUES ($1, $2)`
	execErr := sql.Exec(r.Context(), pool, query, username, body.Content)
	if execErr != nil {
		log.Printf("CreatePostsHandler - Error inserting posts: %v", execErr)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"Post created successfully"}`))
}
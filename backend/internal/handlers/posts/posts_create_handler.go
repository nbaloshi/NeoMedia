package posts

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type PostCreateBody struct {
	UserId  string `json:"userId"`
	Content string `json:"content"`
}

func PostsCreateHandler(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	var body PostCreateBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("CreatePostsHandler - Error decoding body: %v", err)
		utils.RespondError(w, http.StatusBadRequest, "Bad request")
		return
	}

	query := `INSERT INTO posts (user_id, content) VALUES ($1, $2)`

	execErr := sql.Exec(r.Context(), conn, query, body.UserId, body.Content)
	if execErr != nil {
		log.Printf("CreatePostsHandler - Error inserting posts: %v", execErr)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"Post created successfully"}`))
}
package comments

import (
	"NeoMedia/db/sql"
	"NeoMedia/internal/sessions"
	"NeoMedia/internal/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type CommentCreateBody struct {
	PostID	string `json:"post_id"`
	Content string `json:"content"`
}

func CommentsCreateHandler(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	userId, ok := sessions.GetUserIDFromContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "Unauthorized")
    	return
	}

	var body CommentCreateBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("CommentsCreateHandler - Error decoding body: %v", err)
		utils.RespondError(w, http.StatusBadRequest, "Bad request")
		return
	}

	var username string
	queryUsername := `SELECT username FROM users WHERE id = $1`
	queryErr := sql.QueryRow(r.Context(), conn, queryUsername, userId).Scan(&username)
	if queryErr != nil {
		log.Printf("CommentsCreateHandler - Error querrying username: %v", queryErr)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	query := `INSERT INTO comments (post_id, username, content) VALUES ($1, $2, $3)`
	execErr := sql.Exec(r.Context(), conn, query, body.PostID, username, body.Content)
	if execErr != nil {
		log.Printf("CommentsCreateHandler - Error inserting comment: %v", execErr)
		utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"Comment created successfully"}`))
}
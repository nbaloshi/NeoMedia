package utils

import "net/http"

func RespondError(w http.ResponseWriter, status int, msg string) {
	http.Error(w, msg, status)
}

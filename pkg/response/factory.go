package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewResponse(
	w http.ResponseWriter,
	success bool,
	message string,
	status int,
) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	r := Response{
		Success: success,
		Message: message,
	}
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println(err)
	}
}

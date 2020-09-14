package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Body    interface{} `json:"body"`
}

func NewJSONResponse(
	w http.ResponseWriter,
	success bool,
	body interface{},
	status int,
) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	r := Response{
		Success: success,
		Body:    body,
	}
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println(err)
	}
}

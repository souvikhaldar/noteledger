package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/souvikhaldar/noteledger/pkg/note"
)

func (s *Server) AddNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Running in AddNote")
		var tn note.TextNote
		if err := json.NewDecoder(r.Body).Decode(&tn); err != nil {
			log.Println("Error in deserializing the payload: ", err)
			http.Error(w, err.Error(), 500)
			return
		}
		log.Printf("Payload: %+v", tn)
		w.Write([]byte("Note added successfully!"))
	}

}

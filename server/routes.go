package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/souvikhaldar/noteledger/pkg/note"
	"github.com/souvikhaldar/noteledger/pkg/response"
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
		userID := 1
		noteType := `text`
		if err := s.NoteSvc.AddNote(userID, noteType, tn.Note, tn.Bucket); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		response.NewJSONResponse(w, true, "Note added successfully", http.StatusOK)
		return
	}
}

func (s *Server) GetNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Running in GetNote")
		vars := mux.Vars(r)
		bucket := vars["bucket"]
		// hardcoding userID and bucket for now
		userID := 1
		//bucket = "tech"
		notes, err := s.NoteSvc.GetNote(userID, bucket)
		if err != nil {
			log.Println(err)
			response.NewJSONResponse(w, false, nil, 500)
			return
		}
		fmt.Println("Notes: ", notes)
		response.NewJSONResponse(w, true, notes, 200)
		return
	}
}

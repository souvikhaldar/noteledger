package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/souvikhaldar/noteledger/pkg/db"
	"github.com/souvikhaldar/noteledger/pkg/note"
)

type Server struct {
	Router  *mux.Router
	DB      *db.DB
	NoteSvc note.Repository
}

func NewServer() *Server {
	db := db.NewDB()
	return &Server{
		Router:  mux.NewRouter(),
		DB:      db,
		NoteSvc: db,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func main() {
	server := NewServer()
	server.Router.HandleFunc("/note", server.AddNote()).Methods("POST")
	server.Router.HandleFunc("/note/{bucket}", server.GetNote()).Methods("GET")
	port := "8192"
	log.Println("Server is running on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, server))
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	return &Server{
		Router: mux.NewRouter(),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func main() {
	server := NewServer()
	server.Router.HandleFunc("/note", server.AddNote()).Methods("POST")
	port := "8192"
	log.Println("Server is running on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, server))
}

package app

import (
	"log"
	"net/http"
	"github.com/google/uuid"
)

type Server struct {
	id string
	mux *http.ServeMux
}

func NewServer(mux *http.ServeMux) *Server {
	return &Server{id: uuid.New().String(), mux: mux}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/", s.getID)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) getID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	// по умолчанию статус 200 Ok
	_, err := w.Write([]byte(s.id))
	if err != nil {
		log.Println(err)
	}
}

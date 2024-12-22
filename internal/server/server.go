package server

import (
	"process_receipts/internal/handlers"
	"net/http"
	"log"
	"database/sql"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(db *sql.DB) *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	s.setupRoutes(db)
	return s
}

func (s *Server) setupRoutes(db *sql.DB) {
	s.router.HandleFunc("/receipts/process", handlers.HandleAddReceipt(db)).Methods("POST")
	s.router.HandleFunc("/receipts/{id}/points", handlers.HandleGetReceiptById(db)).Methods("GET")
}

func (s *Server) Start(addr string) error {
	log.Println("Starting server on port 8080")
	return http.ListenAndServe(addr, s.router)
}

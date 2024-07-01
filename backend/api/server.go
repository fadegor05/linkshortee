package api

import (
	"github.com/fadegor05/linkshortee/backend/services/link"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	linkHandler := link.NewHandler()
	linkHandler.Register(subRouter)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

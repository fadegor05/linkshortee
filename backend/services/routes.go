package services

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) Register(router *mux.Router) {
	router.HandleFunc("/link", h.handleCreateLink).Methods(http.MethodPost)
	router.HandleFunc("/{code}", h.handleGetLink).Methods(http.MethodGet)
}

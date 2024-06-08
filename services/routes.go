package services

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Register the handlers
	router.HandleFunc("/login", h.Login()).Methods("POST")	
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// Implement the login handler
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

}
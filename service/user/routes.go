package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

// New function for the Handler
func NewHandler() *Handler {
	return &Handler{}
}

// Handler Methods
// Routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// Services
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}

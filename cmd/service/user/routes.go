package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

// construtor da clase
func NewHandler() *Handler {
	return &Handler{}
}

// Metodo
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}

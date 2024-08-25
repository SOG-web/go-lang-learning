package user

import "net/http"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /login", h.handleLogin)
	router.HandleFunc("POST /register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {}

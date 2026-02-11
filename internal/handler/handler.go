package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gookit/slog"
	"github.com/kstsm/wb-l4.5/internal/service"
)

type Calculator interface {
	NewRouter() http.Handler
}

type Handler struct {
	service service.Calculator
	log     *slog.Logger
}

func NewHandler(service service.Calculator, log *slog.Logger) Calculator {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/concatenate", h.concatenateHandler)
	r.Mount("/debug/pprof", http.DefaultServeMux)

	return r
}

func (h *Handler) respondJSON(w http.ResponseWriter, result any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(successResponse{Result: result})
	if err != nil {
		h.log.Error("respondJSON errorResponse", "errorResponse", err.Error())
		return
	}
}

func (h *Handler) respondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(errorResponse{Error: message})
	if err != nil {
		h.log.Error("respondError errorResponse", "errorResponse", err.Error())
		return
	}
}

package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kstsm/wb-l4.5/internal/dto"
	"github.com/kstsm/wb-l4.5/internal/service"
)

func (h *Handler) concatenateHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.AddRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	result, err := h.service.Concatenate(r.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrEmptyItems) {
			h.respondError(w, http.StatusBadRequest, "items cannot be empty")
			return
		}
		h.log.Errorf("Service error: %v", err)
		h.respondError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	h.respondJSON(w, result.Result)
}

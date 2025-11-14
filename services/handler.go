package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/guitar-service/crud"
)

type Handler struct {
	service *GuitarService
}

func NewHandler(s *GuitarService) *Handler {
	return &Handler{service: s}
}

// CREATE
func (h *Handler) CreateGuitar(w http.ResponseWriter, r *http.Request) {
	var g crud.GuitarModel
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if err := h.service.CreateGuitar(r.Context(), &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(g)
}

// READ ONE
func (h *Handler) GetGuitar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	guitar, err := h.service.GetGuitar(r.Context(), uint(id))
	if err != nil {
		http.Error(w, "guitar not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(guitar)
}

// READ LIST
func (h *Handler) ListGuitars(w http.ResponseWriter, r *http.Request) {
	guitars, err := h.service.ListGuitars(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(guitars)
}

// UPDATE
func (h *Handler) UpdateGuitar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	var g crud.GuitarModel
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	g.ID = uint(id) // force ID from URL
	if err := h.service.UpdateGuitar(r.Context(), &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(g)
}

// DELETE
func (h *Handler) DeleteGuitar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err := h.service.DeleteGuitar(r.Context(), uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

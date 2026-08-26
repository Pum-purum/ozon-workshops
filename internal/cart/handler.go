package cart

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type addRequest struct {
	Count uint16 `json:"count"`
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("user_id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	sku, err := strconv.ParseInt(r.PathValue("sku_id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid sku_id", http.StatusBadRequest)
		return
	}

	var req addRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	err = h.service.Add(userID, sku, req.Count)
	if errors.Is(err, ErrProductNotFound) {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (h *Handler) Remove(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("user_id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	sku, err := strconv.ParseInt(r.PathValue("sku_id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid sku_id", http.StatusBadRequest)
		return
	}

	h.service.Remove(userID, sku)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Clear(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("user_id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	h.service.Clear(userID)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.PathValue("user_id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	cart := h.service.Get(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

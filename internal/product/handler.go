package product

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(All())
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	skuStr := r.PathValue("id")
	sku, err := strconv.ParseInt(skuStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid sku", http.StatusBadRequest)
		return
	}

	p, found := FindBySKU(sku)
	if !found {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

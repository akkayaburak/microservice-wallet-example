package handler

import (
	"asset-service/internal/model"
	"asset-service/internal/service"
	"encoding/json"
	"net/http"
)

type AssetHandler struct {
	service *service.AssetService
}

func NewAssetHandler(service *service.AssetService) *AssetHandler {
	return &AssetHandler{service: service}
}

func (h *AssetHandler) CreateAsset(w http.ResponseWriter, r *http.Request) {
	var asset model.Asset
	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdAsset, err := h.service.CreateAsset(asset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAsset)
}

func (h *AssetHandler) GetAssets(w http.ResponseWriter, r *http.Request) {
	assets, err := h.service.GetAssets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(assets)
}

func (h *AssetHandler) UpdateAsset(w http.ResponseWriter, r *http.Request) {
	var a *model.Asset
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	createdAsset, err := h.service.UpdateAsset(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdAsset)

}

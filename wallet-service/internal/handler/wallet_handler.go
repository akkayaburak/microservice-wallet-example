package handler

import (
	"encoding/json"
	"net/http"

	"wallet-service/internal/model"
	"wallet-service/internal/service"

	"github.com/gorilla/mux"
)

type WalletHandler struct {
	service *service.WalletService
}

func NewWalletHandler(service *service.WalletService) *WalletHandler {
	return &WalletHandler{service: service}
}

func (h *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var wallet model.Wallet
	if err := json.NewDecoder(r.Body).Decode(&wallet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdWallet, err := h.service.CreateWallet(wallet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdWallet)
}

func (h *WalletHandler) GetWallet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	address := params["address"]
	network := params["network"]

	wallet, err := h.service.GetWallet(address, network)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(wallet)
}

func (h *WalletHandler) DeleteWallet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	address := params["address"]
	network := params["network"]

	err := h.service.DeleteWallet(address, network)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Wallet deleted"))
}

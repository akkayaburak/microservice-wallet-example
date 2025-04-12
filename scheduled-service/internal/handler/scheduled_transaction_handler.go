package handler

import (
	"encoding/json"
	"net/http"
	"scheduled-service/internal/model"
	"scheduled-service/internal/service"
)

type ScheduledTransactionHandler struct {
	service *service.ScheduledTransactionsService
}

func NewAssetHandler(service *service.ScheduledTransactionsService) *ScheduledTransactionHandler {
	return &ScheduledTransactionHandler{service: service}
}

func (h *ScheduledTransactionHandler) CreateScheduledTransaction(w http.ResponseWriter, r *http.Request) {
	var st model.ScheduledTransaction
	if err := json.NewDecoder(r.Body).Decode(&st); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdSt, err := h.service.CreateScheduledTransaction(st)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdSt)
}

func (h *ScheduledTransactionHandler) ListScheduledTransactions(w http.ResponseWriter, r *http.Request) {
	st, err := h.service.ListScheduledTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(st)
}

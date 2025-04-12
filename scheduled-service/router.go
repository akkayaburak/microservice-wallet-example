package router

import (
	"scheduled-service/internal/handler"

	"github.com/gorilla/mux"
)

func SetupRouter(stHandler handler.ScheduledTransactionHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/schedule", stHandler.CreateScheduledTransaction).Methods("POST")
	r.HandleFunc("/schedules", stHandler.ListScheduledTransactions).Methods("GET")
	return r
}

package router

import (
	"wallet-service/internal/handler"

	"github.com/gorilla/mux"
)

func SetupRouter(walletHandler handler.WalletHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/wallet", walletHandler.CreateWallet).Methods("POST")
	r.HandleFunc("/wallet/{address}/{network}", walletHandler.GetWallet).Methods("GET")
	r.HandleFunc("/wallet/{address}/{network}", walletHandler.DeleteWallet).Methods("DELETE")
	return r
}

package router

import (
	"asset-service/internal/handler"

	"github.com/gorilla/mux"
)

func SetupRouter(assetHandler handler.AssetHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/assets", assetHandler.CreateAsset).Methods("POST")
	r.HandleFunc("/assets/list", assetHandler.GetAssets).Methods("GET")
	r.HandleFunc("/assets/update", assetHandler.UpdateAsset).Methods("PUT")

	return r
}

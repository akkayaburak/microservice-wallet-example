package main

import (
	router "asset-service/internal"
	"asset-service/internal/handler"
	"asset-service/internal/repository"
	"asset-service/internal/service"
	"log"
	"net/http"
)

func main() {
	db, err := repository.NewSQLiteDB("assets.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	assetRepo := repository.NewAssetRepository(db)

	assetRepo.InitializeDB()

	assetService := service.NewAssetService(assetRepo)
	assetHandler := handler.NewAssetHandler(assetService)

	r := router.SetupRouter(*assetHandler)

	log.Println("Server is running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}

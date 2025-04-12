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
	// Veritabanı bağlantısını kuruyoruz
	db, err := repository.NewSQLiteDB("assets.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Repository'yi başlatıyoruz
	assetRepo := repository.NewAssetRepository(db)

	assetRepo.InitializeDB()

	// Servis ve handler'ı başlatıyoruz
	assetService := service.NewAssetService(assetRepo)
	assetHandler := handler.NewAssetHandler(assetService)

	// Router'ı oluşturuyoruz
	r := router.SetupRouter(*assetHandler)

	// Sunucuyu başlatıyoruz
	log.Println("Server is running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}

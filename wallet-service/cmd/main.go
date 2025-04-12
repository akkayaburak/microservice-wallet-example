package main

import (
	"log"
	"net/http"

	router "github.com/burak/microservice-example/wallet-service/internal"
	"github.com/burak/microservice-example/wallet-service/internal/handler"
	"github.com/burak/microservice-example/wallet-service/internal/repository"
	"github.com/burak/microservice-example/wallet-service/internal/service"
)

func main() {
	// Veritabanı bağlantısını kuruyoruz
	db, err := repository.NewSQLiteDB("wallets.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Repository'yi başlatıyoruz
	walletRepo := repository.NewWalletRepository(db)

	walletRepo.InitializeDB()

	// Servis ve handler'ı başlatıyoruz
	walletService := service.NewWalletService(walletRepo)
	walletHandler := handler.NewWalletHandler(walletService)

	// Router'ı oluşturuyoruz
	r := router.SetupRouter(*walletHandler)

	// Sunucuyu başlatıyoruz
	log.Println("Server is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

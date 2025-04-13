package main

import (
	"log"
	"net/http"

	router "wallet-service/internal"
	"wallet-service/internal/handler"
	"wallet-service/internal/repository"
	"wallet-service/internal/service"
)

func main() {
	db, err := repository.NewSQLiteDB("wallets.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	walletRepo := repository.NewWalletRepository(db)

	walletRepo.InitializeDB()

	walletService := service.NewWalletService(walletRepo)
	walletHandler := handler.NewWalletHandler(walletService)

	r := router.SetupRouter(*walletHandler)

	log.Println("Server is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

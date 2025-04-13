package main

import (
	"log"
	"net/http"
	router "scheduled-service"
	"scheduled-service/internal/cron"
	"scheduled-service/internal/handler"
	"scheduled-service/internal/repository"
	"scheduled-service/internal/service"
)

func main() {
	db, err := repository.NewSQLiteDB("schedules.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	stRepo := repository.NewScheduledTransactionRepository(db)

	stRepo.InitializeDB()

	stService := service.NewScheduledTransactionService(stRepo)
	stHandler := handler.NewScheduledTransactionHandler(stService)

	go cron.RunScheduledTransactions(db)

	r := router.SetupRouter(*stHandler)

	log.Println("Server is running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}

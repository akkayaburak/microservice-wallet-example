package cron

import (
	"database/sql"
	"log"
	"scheduled-service/internal/model"
	"time"

	cr "github.com/robfig/cron/v3"
)

func RunScheduledTransactions(db *sql.DB) {
	c := cr.New()
	c.AddFunc("@every 1m", func() {
		log.Println("Running scheduled transactions.")

		now := time.Now()

		tx, err := db.Begin()
		if err != nil {
			log.Println("Error beginning transaction:", err)
			return
		}
		defer tx.Rollback()

		rows, err := tx.Query(`SELECT id, wallet_id, to_address, amount, symbol, network, scheduled_at FROM scheduled_transactions WHERE scheduled_at <= ? AND status = 'pending'`, now)
		if err != nil {
			log.Println("Error querying scheduled transactions:", err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var txModel model.ScheduledTransaction
			if err := rows.Scan(&txModel.ID, &txModel.WalletID, &txModel.ToAddress, &txModel.Amount, &txModel.Symbol, &txModel.Network, &txModel.ScheduledAt); err != nil {
				log.Println("Error scanning transaction:", err)
				continue
			}
			_, err = tx.Exec(`UPDATE scheduled_transactions SET status = 'executed' WHERE id = ?`, txModel.ID)
			if err != nil {
				log.Println("Error updating transaction status:", err)
			} else {
				log.Printf("Transaction %d executed successfully", txModel.ID)
			}
		}

		if err := tx.Commit(); err != nil {
			log.Println("Error committing transaction:", err)
			return
		}

		log.Println("Finished running scheduled transactions.")

	})
	c.Start()
}

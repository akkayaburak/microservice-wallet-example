package repository

import (
	"database/sql"
	"scheduled-service/internal/model"
	"testing"
)

func TestCreateScheduledTransaction(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS scheduled_transactions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			wallet_id INTEGER NOT NULL,
			to_address TEXT NOT NULL,
			amount REAL NOT NULL,
			symbol TEXT NOT NULL,
			network TEXT NOT NULL,
			scheduled_at DATETIME NOT NULL,
			status TEXT NOT NULL CHECK(status IN ('pending', 'executed')) DEFAULT 'pending',
			UNIQUE(wallet_id, to_address, scheduled_at)
		);`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	repo := NewScheduledTransactionRepository(db)

	st := model.ScheduledTransaction{WalletID: 3, Symbol: "BTC", Amount: 50, Network: "Bitcoin", Status: "pending"}
	err = repo.CreateScheduledTransaction(st)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

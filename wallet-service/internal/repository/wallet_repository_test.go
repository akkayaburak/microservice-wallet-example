package repository

import (
	"database/sql"
	"testing"
	"wallet-service/internal/model"
)

func TestCreateWallet(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS wallets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			address TEXT NOT NULL,
			network TEXT NOT NULL,
			UNIQUE(address, network)
		);`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	repo := NewWalletRepository(db)

	wallet := model.Wallet{Address: "123sfsdf", Network: "Bitcoin"}
	err = repo.CreateWallet(wallet)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

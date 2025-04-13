package repository

import (
	"asset-service/internal/model"
	"database/sql"
	"testing"
)

func TestCreateAsset(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS assets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			wallet_id INTEGER NOT NULL,
			symbol TEXT NOT NULL,
			amount REAL NOT NULL CHECK(amount >= 0),
			network TEXT NOT NULL,
			UNIQUE(wallet_id, symbol, network)
	)`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	repo := NewAssetRepository(db)

	asset := model.Asset{WalletID: 3, Symbol: "BTC", Amount: 50, Network: "Bitcoin"}
	err = repo.CreateAsset(asset)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestUpdateAsset(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS assets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			wallet_id INTEGER NOT NULL,
			symbol TEXT NOT NULL,
			amount REAL NOT NULL CHECK(amount >= 0),
			network TEXT NOT NULL,
			UNIQUE(wallet_id, symbol, network)
	)`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	repo := NewAssetRepository(db)

	asset := model.Asset{WalletID: 3, Symbol: "BTC", Amount: 50, Network: "Bitcoin"}
	err = repo.CreateAsset(asset)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	asset.Amount = 40
	asset.ID = 1

	_, err = repo.UpdateAsset(&asset)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

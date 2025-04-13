// repository/asset_repository.go
package repository

import (
	"asset-service/internal/model"
	"database/sql"
	"fmt"
	"log"
)

type AssetRepository struct {
	DB *sql.DB
}

func NewAssetRepository(db *sql.DB) *AssetRepository {
	return &AssetRepository{DB: db}
}

func NewSQLiteDB(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (r *AssetRepository) InitializeDB() {
	query := `
	CREATE TABLE IF NOT EXISTS assets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			wallet_id INTEGER NOT NULL,
			symbol TEXT NOT NULL,
			amount REAL NOT NULL CHECK(amount >= 0),
			network TEXT NOT NULL,
			UNIQUE(wallet_id, symbol, network)
		);
	`
	_, err := r.DB.Exec(query)
	if err != nil {
		log.Fatalf("table could not be created: %v", err)
	}
	fmt.Println("assets table is created or already exists.")
}

func (r *AssetRepository) CreateAsset(asset model.Asset) error {
	_, err := r.DB.Exec(`
		INSERT INTO assets(wallet_id, symbol, amount, network)
		VALUES (?, ?, ?, ?)
	`, asset.WalletID, asset.Symbol, asset.Amount, asset.Network)
	if err != nil {
		return fmt.Errorf("could not add the asset: %v", err)
	}
	return nil
}

func (r *AssetRepository) GetAssets() ([]model.Asset, error) {
	rows, err := r.DB.Query(`SELECT id, wallet_id, symbol, amount, network FROM assets`)
	if err != nil {
		return nil, fmt.Errorf("could not get assets: %v", err)
	}
	defer rows.Close()

	var assets []model.Asset
	for rows.Next() {
		var a model.Asset
		if err := rows.Scan(&a.ID, &a.WalletID, &a.Symbol, &a.Amount, &a.Network); err != nil {
			return nil, fmt.Errorf("scan error %v", err)
		}
		assets = append(assets, a)
	}
	return assets, nil
}

func (r *AssetRepository) UpdateAsset(a *model.Asset) (*model.Asset, error) {
	if a.ID == 0 {
		return nil, fmt.Errorf("id is required")
	}

	_, err := r.DB.Exec(`
		UPDATE assets
		SET amount = ?
		WHERE id = ?`,
		a.Amount, a.ID,
	)
	if err != nil {
		return nil, fmt.Errorf("update failed: %v", err)
	}

	return a, err
}

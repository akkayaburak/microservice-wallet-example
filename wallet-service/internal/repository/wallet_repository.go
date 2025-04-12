// repository/wallet_repository.go
package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/burak/microservice-example/wallet-service/internal/model"
)

type WalletRepository struct {
	DB *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{DB: db}
}

func NewSQLiteDB(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, err
	}
	// Veritabanı bağlantısının açıldığını kontrol et
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (r *WalletRepository) InitializeDB() {
	query := `
	CREATE TABLE IF NOT EXISTS wallets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		address TEXT NOT NULL,
		network TEXT NOT NULL,
		UNIQUE(address, network)
	);
	`
	_, err := r.DB.Exec(query)
	if err != nil {
		log.Fatalf("Veritabanı tablosu oluşturulamadı: %v", err)
	}
	fmt.Println("wallets tablosu oluşturuldu veya zaten mevcut.")
}

func (r *WalletRepository) isWalletExists(address, network string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM wallets WHERE address = ? AND network = ?`
	err := r.DB.QueryRow(query, address, network).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *WalletRepository) CreateWallet(wallet model.Wallet) error {
	exists, err := r.isWalletExists(wallet.Address, wallet.Network)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("wallet already exists: same addresss and network combination")
	}

	query := `INSERT INTO wallets (address, network) VALUES (?, ?)`
	_, err = r.DB.Exec(query, wallet.Address, wallet.Network)
	if err != nil {
		return fmt.Errorf("could not create wallet: %w", err)
	}
	return nil
}

func (r *WalletRepository) GetWallet(address, network string) (*model.Wallet, error) {
	query := `SELECT address, network FROM wallets WHERE address = ? AND network = ?`
	row := r.DB.QueryRow(query, address, network)

	var wallet model.Wallet
	if err := row.Scan(&wallet.Address, &wallet.Network); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Eğer wallet bulunamazsa, nil döner
		}
		return nil, fmt.Errorf("could not get wallet: %w", err)
	}
	return &wallet, nil
}

func (r *WalletRepository) DeleteWallet(address, network string) error {
	query := `DELETE FROM wallets WHERE address = ? AND network = ?`
	_, err := r.DB.Exec(query, address, network)
	if err != nil {
		return fmt.Errorf("could not delete wallet: %w", err)
	}
	return nil
}

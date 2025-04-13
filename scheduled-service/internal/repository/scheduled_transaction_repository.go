package repository

import (
	"database/sql"
	"fmt"
	"log"
	"scheduled-service/internal/model"
)

type ScheduledTransactionRepository struct {
	DB *sql.DB
}

func NewScheduledTransactionRepository(db *sql.DB) *ScheduledTransactionRepository {
	return &ScheduledTransactionRepository{DB: db}
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

func (r *ScheduledTransactionRepository) InitializeDB() {
	query := `
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
		);
	`
	_, err := r.DB.Exec(query)
	if err != nil {
		log.Fatalf("table could not be created: %v", err)
	}
	fmt.Println("scheduled_transactions table is created or already exists")
}

func (r *ScheduledTransactionRepository) CreateScheduledTransaction(st model.ScheduledTransaction) error {
	_, err := r.DB.Exec(`
			INSERT INTO scheduled_transactions(wallet_id, to_address, amount, symbol, network, scheduled_at, status)
			VALUES (?, ?, ?, ?, ?, ?, ?)`,
		st.WalletID, st.ToAddress, st.Amount, st.Symbol, st.Network, st.ScheduledAt, st.Status,
	)
	if err != nil {
		return fmt.Errorf("insert failed: %v", err)
	}

	return nil
}

func (r *ScheduledTransactionRepository) ListScheduledTransactions() ([]model.ScheduledTransaction, error) {
	rows, err := r.DB.Query(`SELECT id, wallet_id, to_address, amount, symbol, network, scheduled_at, status FROM scheduled_transactions`)
	if err != nil {
		return nil, fmt.Errorf("db error: %v", err)
	}
	defer rows.Close()

	var transactions []model.ScheduledTransaction
	for rows.Next() {
		var st model.ScheduledTransaction
		if err := rows.Scan(&st.ID, &st.WalletID, &st.ToAddress, &st.Amount, &st.Symbol, &st.Network, &st.ScheduledAt, &st.Status); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		transactions = append(transactions, st)
	}

	return transactions, nil
}

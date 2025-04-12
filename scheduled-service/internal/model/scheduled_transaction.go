package model

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type ScheduledTransaction struct {
	ID          int       `json:"id"`
	WalletID    int       `json:"wallet_id"`
	ToAddress   string    `json:"to_address"`
	Amount      float64   `json:"amount"`
	Symbol      string    `json:"symbol"`
	Network     string    `json:"network"`
	ScheduledAt time.Time `json:"scheduled_at"`
	Status      string    `json:"status"` // pending, executed
}

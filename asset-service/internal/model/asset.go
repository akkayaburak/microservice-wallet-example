package model

import (
	_ "github.com/mattn/go-sqlite3"
)

type Asset struct {
	ID       int     `json:"id"`
	WalletID int     `json:"wallet_id"`
	Symbol   string  `json:"symbol"` // BTC, ETH, USDT...
	Amount   float64 `json:"amount"`
	Network  string  `json:"network"` // Bitcoin, Ethereum, vs.
}

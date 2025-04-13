package model

import (
	_ "github.com/mattn/go-sqlite3"
)

type Wallet struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
	Network string `json:"network"`
}

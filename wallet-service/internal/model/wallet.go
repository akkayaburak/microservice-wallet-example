package model

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./wallets.db")
	if err != nil {
		panic(err)
	}
}

func CreateWallet(wallet Wallet) error {
	query := "INSERT INTO wallets (address, network) VALUES (?, ?)"
	_, err := db.Exec(query, wallet.Address, wallet.Network)
	return err
}

func GetWallet(address, network string) (Wallet, error) {
	var wallet Wallet
	query := "SELECT id, address, network FROM wallets WHERE address = ? AND network = ?"
	err := db.QueryRow(query, address, network).Scan(&wallet.ID, &wallet.Address, &wallet.Network)
	if err != nil {
		if err == sql.ErrNoRows {
			return Wallet{}, errors.New("wallet not found")
		}
		return Wallet{}, err
	}
	return wallet, nil
}

func DeleteWallet(address, network string) error {
	query := "DELETE FROM wallets WHERE address = ? AND network = ?"
	_, err := db.Exec(query, address, network)
	return err
}

type Wallet struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
	Network string `json:"network"`
}

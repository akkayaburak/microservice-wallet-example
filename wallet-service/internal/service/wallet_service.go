package service

import (
	"wallet-service/internal/model"
	"wallet-service/internal/repository"
)

type WalletService struct {
	repo *repository.WalletRepository
}

func NewWalletService(repo *repository.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) CreateWallet(wallet model.Wallet) (model.Wallet, error) {
	err := s.repo.CreateWallet(wallet)
	if err != nil {
		return model.Wallet{}, err
	}
	return wallet, nil
}

func (s *WalletService) GetWallet(address, network string) (model.Wallet, error) {
	wallet, err := s.repo.GetWallet(address, network)
	if err != nil {
		return model.Wallet{}, err
	}
	return *wallet, nil
}

func (s *WalletService) DeleteWallet(address, network string) error {
	err := s.repo.DeleteWallet(address, network)
	if err != nil {
		return err
	}
	return nil
}

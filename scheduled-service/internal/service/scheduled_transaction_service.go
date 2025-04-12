package service

import (
	"scheduled-service/internal/model"
	"scheduled-service/internal/repository"
)

type ScheduledTransactionsService struct {
	repo *repository.ScheduledTransactionRepository
}

func NewAssetService(repo *repository.ScheduledTransactionRepository) *ScheduledTransactionsService {
	return &ScheduledTransactionsService{repo: repo}
}

func (s *ScheduledTransactionsService) CreateScheduledTransaction(st model.ScheduledTransaction) (model.ScheduledTransaction, error) {
	err := s.repo.CreateScheduledTransaction(st)
	if err != nil {
		return model.ScheduledTransaction{}, err
	}
	return st, nil
}

func (s *ScheduledTransactionsService) ListScheduledTransactions() ([]model.ScheduledTransaction, error) {
	st, err := s.repo.ListScheduledTransactions()
	if err != nil {
		return []model.ScheduledTransaction{}, err
	}
	return st, nil
}

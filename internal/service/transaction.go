package service

import (
	"market/internal/repository"
	"market/models"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) GetInfo(ID int) (models.Info, error) {
	return s.repo.GetInfo(ID)

}

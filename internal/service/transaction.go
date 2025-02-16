package service

import (
	"errors"
	"fmt"
	"market/internal/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) SendCoins(fromUserID, toUserID, amount int) error {
	if fromUserID == toUserID {
		return errors.New("sender and receiver cannot be the same user")
	}

	coins, err := s.repo.GetCoinsByID(fromUserID)
	if err != nil {
		return fmt.Errorf("failed to get balance for user %d: %w", fromUserID, err)
	}

	if amount > coins {
		return errors.New("not enough coins")
	}

	if err := s.repo.SendCoins(fromUserID, toUserID, amount); err != nil {
		return fmt.Errorf("failed to complete transaction from %d to %d: %w", fromUserID, toUserID, err)
	}

	return nil
}

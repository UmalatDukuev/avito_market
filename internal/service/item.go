package service

import (
	"errors"
	"fmt"
	"market/internal/repository"
)

type ItemService struct {
	repo repository.Item
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) BuyItem(userID int, itemName string) error {
	coins, err := s.repo.GetCoinsByID(userID)
	if err != nil {
		return fmt.Errorf("failed to get balance for user %d: %w", userID, err)
	}

	itemPrice, err := s.repo.GetItemPrice(itemName)
	if err != nil {
		return fmt.Errorf("failed to get price of item %s: %w", itemName, err)
	}

	if itemPrice > coins {
		return errors.New("not enough coins")
	}

	return s.repo.BuyItem(userID, itemName)
}

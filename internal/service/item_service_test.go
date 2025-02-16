package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockItemRepo struct {
	mock.Mock
}

func (m *MockItemRepo) GetCoinsByID(userID int) (int, error) {
	args := m.Called(userID)
	return args.Int(0), args.Error(1)
}

func (m *MockItemRepo) GetItemPrice(itemName string) (int, error) {
	args := m.Called(itemName)
	return args.Int(0), args.Error(1)
}

func (m *MockItemRepo) BuyItem(userID int, itemName string) error {
	args := m.Called(userID, itemName)
	return args.Error(0)
}

func TestBuyItem_Success(t *testing.T) {
	mockRepo := new(MockItemRepo)
	service := &ItemService{repo: mockRepo}

	userID := 1
	itemName := "cup"
	mockRepo.On("GetCoinsByID", userID).Return(50, nil)
	mockRepo.On("GetItemPrice", itemName).Return(20, nil)
	mockRepo.On("BuyItem", userID, itemName).Return(nil)

	err := service.BuyItem(userID, itemName)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBuyItem_NotEnoughCoins(t *testing.T) {
	mockRepo := new(MockItemRepo)
	service := &ItemService{repo: mockRepo}

	userID := 1
	itemName := "cup"
	mockRepo.On("GetCoinsByID", userID).Return(10, nil)
	mockRepo.On("GetItemPrice", itemName).Return(20, nil)

	err := service.BuyItem(userID, itemName)

	assert.Error(t, err)
	assert.Equal(t, "not enough coins", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestBuyItem_ItemNotFound(t *testing.T) {
	mockRepo := new(MockItemRepo)
	service := &ItemService{repo: mockRepo}

	userID := 1
	itemName := "unknown"
	mockRepo.On("GetCoinsByID", userID).Return(100, nil)
	mockRepo.On("GetItemPrice", itemName).Return(0, errors.New("item not found"))

	err := service.BuyItem(userID, itemName)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "item not found")
	mockRepo.AssertExpectations(t)
}

func TestBuyItem_GetCoinsError(t *testing.T) {
	mockRepo := new(MockItemRepo)
	service := &ItemService{repo: mockRepo}

	userID := 1
	itemName := "cup"

	mockRepo.On("GetCoinsByID", userID).Return(0, errors.New("database error"))

	err := service.BuyItem(userID, itemName)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get balance for user")
}

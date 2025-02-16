package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepo struct {
	mock.Mock
}

func (m *MockTransactionRepo) GetCoinsByID(userID int) (int, error) {
	args := m.Called(userID)
	return args.Int(0), args.Error(1)
}

func (m *MockTransactionRepo) SendCoins(fromUserID, toUserID, amount int) error {
	args := m.Called(fromUserID, toUserID, amount)
	return args.Error(0)
}

func TestSendCoins_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	service := &TransactionService{repo: mockRepo}

	fromUser := 1
	toUser := 2
	amount := 20

	mockRepo.On("GetCoinsByID", fromUser).Return(50, nil)
	mockRepo.On("SendCoins", fromUser, toUser, amount).Return(nil)

	err := service.SendCoins(fromUser, toUser, amount)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSendCoins_SameUser(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	service := &TransactionService{repo: mockRepo}

	err := service.SendCoins(1, 1, 10)

	assert.Error(t, err)
	assert.Equal(t, "sender and receiver cannot be the same user", err.Error())
}

func TestSendCoins_NotEnoughCoins(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	service := &TransactionService{repo: mockRepo}

	fromUser := 1
	toUser := 2
	amount := 50

	mockRepo.On("GetCoinsByID", fromUser).Return(20, nil)

	err := service.SendCoins(fromUser, toUser, amount)

	assert.Error(t, err)
	assert.Equal(t, "not enough coins", err.Error())
}

func TestSendCoins_ErrorInRepo(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	service := &TransactionService{repo: mockRepo}

	fromUser := 1
	toUser := 2
	amount := 20

	mockRepo.On("GetCoinsByID", fromUser).Return(50, nil)
	mockRepo.On("SendCoins", fromUser, toUser, amount).Return(errors.New("db error"))

	err := service.SendCoins(fromUser, toUser, amount)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to complete transaction")
}

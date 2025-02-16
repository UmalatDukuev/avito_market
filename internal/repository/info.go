package repository

import (
	"fmt"
	"market/models"

	"github.com/jmoiron/sqlx"
)

type InfoPostgres struct {
	db *sqlx.DB
}

func NewInfoPostgres(db *sqlx.DB) *InfoPostgres {
	return &InfoPostgres{db: db}
}

func (r *InfoPostgres) GetInfo(userID int) (models.Info, error) {
	var info models.Info

	query := "SELECT coins FROM users WHERE id = $1"
	err := r.db.QueryRow(query, userID).Scan(&info.Coins)
	if err != nil {
		return info, fmt.Errorf("failed to get user coins: %w", err)
	}

	inventoryQuery := `
		SELECT m.id, m.name, m.price, i.quantity 
		FROM inventory i
		JOIN merch m ON i.merch_id = m.id
		WHERE i.user_id = $1`
	rows, err := r.db.Query(inventoryQuery, userID)
	if err != nil {
		return info, fmt.Errorf("failed to get user inventory: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.Quantity)
		if err != nil {
			return info, fmt.Errorf("failed to scan inventory item: %w", err)
		}
		info.Inventory = append(info.Inventory, item)
	}

	receivedQuery := `
		SELECT from_user_id, amount 
		FROM transactions 
		WHERE to_user_id = $1`
	receivedRows, err := r.db.Query(receivedQuery, userID)
	if err != nil {
		return info, fmt.Errorf("failed to get received transactions: %w", err)
	}
	defer receivedRows.Close()

	for receivedRows.Next() {
		var received models.ReceivedCoinRequest
		err := receivedRows.Scan(&received.FromUser, &received.Amount)
		if err != nil {
			return info, fmt.Errorf("failed to scan received transaction: %w", err)
		}
		info.CoinHistory.Received = append(info.CoinHistory.Received, received)
	}

	sentQuery := `
		SELECT to_user_id, amount 
		FROM transactions 
		WHERE from_user_id = $1`
	sentRows, err := r.db.Query(sentQuery, userID)
	if err != nil {
		return info, fmt.Errorf("failed to get sent transactions: %w", err)
	}
	defer sentRows.Close()

	for sentRows.Next() {
		var sent models.SendCoinRequest
		err := sentRows.Scan(&sent.ToUser, &sent.Amount)
		if err != nil {
			return info, fmt.Errorf("failed to scan sent transaction: %w", err)
		}
		info.CoinHistory.Sent = append(info.CoinHistory.Sent, sent)
	}

	return info, nil
}

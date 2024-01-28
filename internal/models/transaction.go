package models

import (
	"context"
	"database/sql"
	"time"
)

type Transaction struct {
	ContactNumber     string  `json:"contact_number"`
	AmountAdminFee    float64 `json:"amount_admin_fee"`
	AmountOTR         float64 `json:"amount_otr"`
	AmountInstallment float64 `json:"amount_installment"`
	AmountInterest    float64 `json:"amount_interest"`

	UserID      int    `json:"user_id"`
	ProductID   int    `json:"product_id"`
	ProductName string `json:"asset_name"`
}

type TransactionModel struct {
	DB *sql.DB
}

func (m TransactionModel) Trx(contactNumber string, user *User, limit *Limit, product *Product) (int, error) {
	if limit.UserID != user.ID {
		return 0, ErrNoMatched
	}
	if !limit.IsEligible(product.Price) {
		return 0, ErrLimitInsufficient
	}

	tx, err := m.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// Transaction: New
	query := `
		INSERT INTO transactions
		    (contact_number, amount_admin_fee, amount_otr, amount_installment, amount_interest, product_id, asset_name, user_id)
		VALUES
		    (?, ?, ?, ?, ?, ?, ?, ?)
	`
	args := []any{
		contactNumber,
		limit.AmountAdminFee,
		product.Price,
		limit.PayPerMonth * float64(limit.Month),
		limit.RatePerMonth * float64(limit.Month),
		product.ID,
		product.Name,
		user.ID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	// Limit: reduce
	query = `UPDATE limits SET consumer_limit = consumer_limit - ? WHERE id = ?`
	args = []any{product.Price, limit.ID}
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

func (m TransactionModel) Get(id int) (*Transaction, error) {
	query := `
		SELECT contact_number, amount_admin_fee, amount_otr, amount_installment, amount_interest, product_id, asset_name, user_id
		FROM transactions
		WHERE id = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var transaction Transaction
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&transaction.ContactNumber,
		&transaction.AmountAdminFee,
		&transaction.AmountOTR,
		&transaction.AmountInstallment,
		&transaction.AmountInterest,
		&transaction.ProductID,
		&transaction.ProductName,
		&transaction.UserID,
	)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

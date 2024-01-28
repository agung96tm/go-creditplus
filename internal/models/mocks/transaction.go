package mocks

import (
	"github.com/agung96tm/go-creditplus/internal/models"
)

type TransactionModel struct {
}

func (m TransactionModel) Trx(contactNumber string, user *models.User, limit *models.Limit, product *models.Product) (int, error) {
	return 0, nil
}

func (m TransactionModel) Get(id int) (*models.Transaction, error) {
	return &models.Transaction{}, nil
}

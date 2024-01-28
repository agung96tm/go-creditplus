package mocks

import "github.com/agung96tm/go-creditplus/internal/models"

func NewModel() *models.Models {
	return &models.Models{
		User:        UserModel{},
		Limit:       LimitModel{},
		Product:     ProductModel{},
		Config:      ConfigModel{},
		ConfigRate:  ConfigRateModel{},
		Transaction: TransactionModel{},
	}
}

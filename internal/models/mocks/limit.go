package mocks

import (
	"github.com/agung96tm/go-creditplus/internal/models"
)

type LimitModel struct {
}

func (m LimitModel) GetLimitByConsumer(user *models.User) ([]*models.Limit, error) {
	return make([]*models.Limit, 0), nil
}

func (m LimitModel) GetLimitsByUserAndProduct(user *models.User, product *models.Product) ([]*models.Limit, error) {
	return make([]*models.Limit, 0), nil
}

func (m LimitModel) ReduceLimit(limit *models.Limit, amount float64) (bool, error) {
	return false, nil
}

func (m LimitModel) GetWithProductCalc(id int, product *models.Product) (*models.Limit, error) {
	return &models.Limit{}, nil
}

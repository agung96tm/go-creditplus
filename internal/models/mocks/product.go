package mocks

import (
	"github.com/agung96tm/go-creditplus/internal/models"
)

type ProductModel struct {
}

func (m ProductModel) GetAll() ([]*models.Product, error) {
	return make([]*models.Product, 0), nil
}

func (m ProductModel) Get(id int) (*models.Product, error) {
	return &models.Product{}, nil
}

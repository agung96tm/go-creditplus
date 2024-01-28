package mocks

import (
	"github.com/agung96tm/go-creditplus/internal/models"
)

type ProductModel struct {
}

func (m ProductModel) GetAll() ([]*models.Product, error) {
	return []*models.Product{
		{
			ID:          1,
			Name:        "Product1",
			Price:       200000.0,
			Description: "Product 1",
			PartnerID:   1,
			PartnerName: "My Partner",
		},
	}, nil
}

func (m ProductModel) Get(id int) (*models.Product, error) {
	if id != 1 {
		return nil, models.ErrNoDataFound
	}

	product := models.Product{
		ID:          1,
		Name:        "Product1",
		Price:       200000.0,
		Description: "Product 1",
		PartnerID:   1,
		PartnerName: "My Partner",
	}

	return &product, nil
}

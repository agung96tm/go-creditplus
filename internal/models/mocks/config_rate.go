package mocks

import (
	"github.com/agung96tm/go-creditplus/internal/models"
)

type ConfigRateModel struct {
}

func (m ConfigRateModel) GetRateByMonth(month int) (*models.ConfigRate, error) {
	return &models.ConfigRate{}, nil
}

package mocks

import (
	"github.com/agung96tm/go-creditplus/internal/models"
)

type ConfigModel struct {
}

func (m ConfigModel) Get() (*models.Config, error) {
	return &models.Config{}, nil
}

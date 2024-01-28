package mocks

import (
	"github.com/agung96tm/go-creditplus/internal/models"
)

type UserModel struct {
}

func (m UserModel) GetById(id int) (*models.User, error) {
	return &models.User{}, nil
}

func (m UserModel) GetAll() ([]*models.User, error) {
	return make([]*models.User, 0), nil
}

func (m UserModel) Exists(id int) (bool, error) {
	return false, nil
}

func (m UserModel) Authenticate(nik, password string) (int, error) {
	return 0, nil
}

func (m UserModel) UpdatePassword(id int, password string) error {
	return nil
}

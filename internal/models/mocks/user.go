package mocks

import (
	"github.com/agung96tm/go-creditplus/internal/models"
)

type UserModel struct {
}

func (m UserModel) GetById(id int) (*models.User, error) {
	if id != 1 {
		return nil, models.ErrInvalidCredentials
	}
	return &models.User{
		ID:          1,
		NIK:         "123456789",
		FullName:    "Example Fullname",
		LegalName:   "Example LegalName",
		PlaceBirth:  "Banten",
		DateBirth:   "",
		Salary:      28888.0,
		IDCardPhoto: "https://example.com/1.jpg",
		SelfiePhoto: "https://example.com/1.jpg",
	}, nil
}

func (m UserModel) GetAll() ([]*models.User, error) {
	return make([]*models.User, 0), nil
}

func (m UserModel) Exists(id int) (bool, error) {
	if id != 1 {
		return false, nil
	}
	return true, nil
}

func (m UserModel) Authenticate(nik, password string) (int, error) {
	if nik == "123456789" && password == "pa$$word" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func (m UserModel) UpdatePassword(id int, password string) error {
	if id != 1 {
		return nil
	}
	return models.ErrNoDataFound
}

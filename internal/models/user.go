package models

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID          int     `json:"id"`
	NIK         string  `json:"nik"`
	FullName    string  `json:"full_name"`
	LegalName   string  `json:"legal_name"`
	PlaceBirth  string  `json:"place_birth"`
	DateBirth   string  `json:"date_birth"`
	Salary      float64 `json:"salary"`
	IDCardPhoto string  `json:"id_card_photo"`
	SelfiePhoto string  `json:"selfie_photo"`
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) GetById(id int) (*User, error) {
	query := `
		SELECT id, nik, full_name, legal_name, place_birth, date_birth, salary, id_card_photo, selfie_photo
		FROM consumers
		WHERE id = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var consumer User
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&consumer.ID,
		&consumer.NIK,
		&consumer.FullName,
		&consumer.LegalName,
		&consumer.PlaceBirth,
		&consumer.DateBirth,
		&consumer.Salary,
		&consumer.IDCardPhoto,
		&consumer.SelfiePhoto,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoDataFound
		default:
			return nil, err
		}
	}
	return &consumer, nil
}

func (m UserModel) GetAll() ([]*User, error) {
	query := `
		SELECT id, nik, full_name, legal_name, place_birth, date_birth, salary, id_card_photo, selfie_photo
		FROM consumers
		ORDER BY id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoDataFound
		default:
			return nil, err
		}
	}
	defer rows.Close()

	var consumers []*User
	for rows.Next() {
		var consumer User
		err := rows.Scan(
			&consumer.ID,
			&consumer.NIK,
			&consumer.FullName,
			&consumer.LegalName,
			&consumer.PlaceBirth,
			&consumer.DateBirth,
			&consumer.Salary,
			&consumer.IDCardPhoto,
			&consumer.SelfiePhoto)
		if err != nil {
			return nil, err
		}

		consumers = append(consumers, &consumer)
	}
	return consumers, nil
}

func (m UserModel) Exists(id int) (bool, error) {
	return false, nil
}

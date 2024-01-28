package models

import (
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
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

type UserModelInterface interface {
	GetById(id int) (*User, error)
	GetAll() ([]*User, error)
	Exists(id int) (bool, error)
	Authenticate(nik, password string) (int, error)
	UpdatePassword(id int, password string) error
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) GetById(id int) (*User, error) {
	query := `
		SELECT id, nik, full_name, legal_name, place_birth, date_birth, salary, id_card_photo, selfie_photo
		FROM users
		WHERE id = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.NIK,
		&user.FullName,
		&user.LegalName,
		&user.PlaceBirth,
		&user.DateBirth,
		&user.Salary,
		&user.IDCardPhoto,
		&user.SelfiePhoto,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoDataFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (m UserModel) GetAll() ([]*User, error) {
	query := `
		SELECT id, nik, full_name, legal_name, place_birth, date_birth, salary, id_card_photo, selfie_photo
		FROM users
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

	var users []*User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.NIK,
			&user.FullName,
			&user.LegalName,
			&user.PlaceBirth,
			&user.DateBirth,
			&user.Salary,
			&user.IDCardPhoto,
			&user.SelfiePhoto)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}
	return users, nil
}

func (m UserModel) Exists(id int) (bool, error) {
	var exists bool
	stmt := `SELECT EXISTS(SELECT true FROM users WHERE id = ?)`
	err := m.DB.QueryRow(stmt, id).Scan(&exists)
	return exists, err
}

func (m UserModel) Authenticate(nik, password string) (int, error) {
	var id int
	var savedPassword string

	stmt := `SELECT id, password FROM users WHERE nik = ?`
	err := m.DB.QueryRow(stmt, nik).Scan(&id, &savedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (m UserModel) UpdatePassword(id int, password string) error {
	stmt := `
		UPDATE users
		SET password = ?
		WHERE id = ?
	`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	_, err = m.DB.Exec(stmt, id, string(hashedPassword))
	if err != nil {
		return err
	}
	return nil
}

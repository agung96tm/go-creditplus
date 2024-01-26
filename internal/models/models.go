package models

import (
	"database/sql"
	"errors"
)

var ErrNoDataFound = errors.New("data not found")
var ErrDuplicateEmail = errors.New("duplicate email")

type Models struct {
	User  UserModel
	Limit LimitModel
}

func New(db *sql.DB) *Models {
	return &Models{
		UserModel{DB: db},
		LimitModel{DB: db},
	}
}

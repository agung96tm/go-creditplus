package models

import (
	"database/sql"
	"errors"
)

var NoDataFound = errors.New("data not found")
var ErrDuplicateEmail = errors.New("duplicate email")

type Models struct {
	User UserModel
}

func New(db *sql.DB) *Models {
	return &Models{
		UserModel{DB: db},
	}
}

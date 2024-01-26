package models

import (
	"database/sql"
)

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

package models

import (
	"database/sql"
)

type Models struct {
	User    UserModel
	Limit   LimitModel
	Product ProductModel
}

func New(db *sql.DB) *Models {
	return &Models{
		UserModel{DB: db},
		LimitModel{DB: db},
		ProductModel{DB: db},
	}
}

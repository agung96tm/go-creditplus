package models

import (
	"database/sql"
)

type Models struct {
	User       UserModel
	Limit      LimitModel
	Product    ProductModel
	Config     ConfigModel
	ConfigRate ConfigRateModel
}

func New(db *sql.DB) *Models {
	return &Models{
		UserModel{DB: db},
		LimitModel{DB: db},
		ProductModel{DB: db},
		ConfigModel{DB: db},
		ConfigRateModel{DB: db},
	}
}

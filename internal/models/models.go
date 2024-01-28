package models

import (
	"database/sql"
)

type Models struct {
	User        UserModelInterface
	Limit       LimitModelInterface
	Product     ProductModelInterface
	Config      ConfigModelInterface
	ConfigRate  ConfigRateInterface
	Transaction TransactionModelInterface
}

func New(db *sql.DB) *Models {
	return &Models{
		User:        UserModel{DB: db},
		Limit:       LimitModel{DB: db},
		Product:     ProductModel{DB: db},
		Config:      ConfigModel{DB: db},
		ConfigRate:  ConfigRateModel{DB: db},
		Transaction: TransactionModel{DB: db},
	}
}

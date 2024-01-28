package models

import (
	"context"
	"database/sql"
	"time"
)

type Config struct {
	ID       int     `json:"id"`
	AdminFee float64 `json:"admin_fee"`
}

type ConfigModelInterface interface {
	Get() (*Config, error)
}

type ConfigModel struct {
	DB *sql.DB
}

func (m ConfigModel) Get() (*Config, error) {
	query := `SELECT id, admin_fee FROM configs LIMIT 1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var config Config
	err := m.DB.QueryRowContext(ctx, query).Scan(&config.ID, &config.AdminFee)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

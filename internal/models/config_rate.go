package models

import (
	"context"
	"database/sql"
	"time"
)

type ConfigRate struct {
	ID         int     `json:"id"`
	Month      int     `json:"month"`
	Percentage float64 `json:"percentage"`
}

type ConfigRateModel struct {
	DB *sql.DB
}

func (m ConfigRateModel) GetRateByMonth(month int) (*ConfigRate, error) {
	query := `SELECT id, month, percentage FROM config_rates WHERE month = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rate ConfigRate
	err := m.DB.QueryRowContext(ctx, query, month).Scan(&rate.ID, &rate.Month, &rate.Percentage)
	if err != nil {
		return nil, err
	}

	return &rate, nil
}

package models

import (
	"context"
	"database/sql"
	"time"
)

type Limit struct {
	ID            int     `json:"id"`
	ConsumerID    int     `json:"consumer_id"`
	Month         int     `json:"month"`
	ConsumerLimit float64 `json:"consumer_limit"`
}

type LimitModel struct {
	DB *sql.DB
}

func (m LimitModel) GetLimitByConsumer(consumer *Consumer) ([]*Limit, error) {
	query := `
			SELECT id, consumer_id, month, consumer_limit
			FROM limits
			WHERE consumer_id = ?
			ORDER BY month
		`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, consumer.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var limits []*Limit
	for rows.Next() {
		var limit Limit
		err := rows.Scan(
			&limit.ID,
			&limit.ConsumerID,
			&limit.Month,
			&limit.ConsumerLimit,
		)
		if err != nil {
			return nil, err
		}

		limits = append(limits, &limit)
	}

	return limits, nil
}

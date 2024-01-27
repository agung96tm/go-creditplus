package models

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Limit struct {
	ID            int     `json:"id"`
	UserID        int     `json:"consumer_id"`
	Month         int     `json:"month"`
	ConsumerLimit float64 `json:"consumer_limit"`

	Total          float64 `json:"total"`
	AdminFee       float64 `json:"admin_fee"`
	PayPerMonth    float64 `json:"pay_per_month"`
	PercentageRate float64 `json:"percentage_rate"`
}

type LimitModel struct {
	DB *sql.DB
}

func (m LimitModel) GetLimitByConsumer(user *User) ([]*Limit, error) {
	query := `
			SELECT id, user_id, month, consumer_limit
			FROM limits
			WHERE user_id = ?
			ORDER BY month
		`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var limits []*Limit
	for rows.Next() {
		var limit Limit
		err := rows.Scan(
			&limit.ID,
			&limit.UserID,
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

func (m LimitModel) GetLimitsByUserAndProduct(user *User, product *Product) ([]*Limit, error) {
	query := `
		SELECT
			id,
			user_id,
			month,
			consumer_limit,
			admin_fee,
			total,
			(total / l.month + (total / l.month * percentage / 100)) AS pay_per_month,
			percentage
		FROM (
				 SELECT
					 l.id,
					 l.user_id,
					 l.month,
					 l.consumer_limit,
					 c.admin_fee,
					 ? + ? * c.admin_fee / 100 AS total,
					 cr.percentage
				 FROM
					 limits l
						 JOIN
					 config_rates cr ON cr.month = l.month
						 CROSS JOIN
					 (SELECT admin_fee FROM configs LIMIT 1) c
				 WHERE
					 l.user_id = ?
				   AND l.consumer_limit >= ?
			 ) AS l
		ORDER BY
			month;
		`
	args := []any{product.Price, product.Price, user.ID, product.Price}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoDataFound
		default:
			return nil, err
		}
	}

	var limits []*Limit
	for rows.Next() {
		var limit Limit
		err := rows.Scan(
			&limit.ID,
			&limit.UserID,
			&limit.Month,
			&limit.ConsumerLimit,
			&limit.AdminFee,
			&limit.Total,
			&limit.PayPerMonth,
			&limit.PercentageRate,
		)
		if err != nil {
			return nil, err
		}

		limits = append(limits, &limit)
	}

	return limits, nil
}

func (m LimitModel) Get(id int) (*Limit, error) {
	query := `
			SELECT id, user_id, month, consumer_limit
			FROM limits
			WHERE id = ?
			ORDER BY month
		`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var limit Limit
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&limit.ID,
		&limit.UserID,
		&limit.Month,
		&limit.ConsumerLimit,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoDataFound
		default:
			return nil, err
		}
	}

	return &limit, nil
}

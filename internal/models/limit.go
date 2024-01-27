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
	AmountAdminFee float64 `json:"amount_admin_fee"`
	PayPerMonth    float64 `json:"pay_per_month"`
	RatePerMonth   float64 `json:"rate_per_month"`
	PercentageRate float64 `json:"percentage_rate"`
}

func (l Limit) IsEligible(price float64) bool {
	return l.ConsumerLimit >= price
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
		    amount_admin_fee,
			total,
		    (total / l.month * percentage / 100) AS rate_per_month,
			(total / l.month + (total / l.month * percentage / 100)) AS pay_per_month,
			percentage
		FROM (
				 SELECT
					 l.id,
					 l.user_id,
					 l.month,
					 l.consumer_limit,
					 c.admin_fee,
		             ? * c.admin_fee / 100 AS amount_admin_fee,
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
	args := []any{product.Price, product.Price, product.Price, user.ID, product.Price}

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
			&limit.AmountAdminFee,
			&limit.Total,
			&limit.RatePerMonth,
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

func (m LimitModel) ReduceLimit(limit *Limit, amount float64) (bool, error) {
	query := `UPDATE limits SET consumer_limit = consumer_limit - ? WHERE id = ?`
	args := []any{amount, limit.ID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r, err := m.DB.QueryContext(ctx, query, args...)
	defer r.Close()

	if err != nil {
		return false, err
	}
	return true, nil
}

func (m LimitModel) GetWithProductCalc(id int, product *Product) (*Limit, error) {
	query := `
		SELECT
			id,
			user_id,
			month,
			consumer_limit,
			admin_fee,
		    amount_admin_fee,
			total,
		    (total / l.month * percentage / 100) AS rate_per_month,
			(total / l.month + (total / l.month * percentage / 100)) AS pay_per_month,
			percentage
		FROM (
				 SELECT
					 l.id,
					 l.user_id,
					 l.month,
					 l.consumer_limit,
					 c.admin_fee,
		             ? * c.admin_fee / 100 AS amount_admin_fee,
					 ? + ? * c.admin_fee / 100 AS total,
					 cr.percentage
				 FROM
					 limits l
						 JOIN
					 config_rates cr ON cr.month = l.month
						 CROSS JOIN
					 (SELECT admin_fee FROM configs LIMIT 1) c
				 WHERE
					 l.id = ?
				   AND l.consumer_limit >= ?
			 ) AS l
		ORDER BY
			month;
		`
	args := []any{product.Price, product.Price, product.Price, id, product.Price}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var limit Limit
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&limit.ID,
		&limit.UserID,
		&limit.Month,
		&limit.ConsumerLimit,
		&limit.AdminFee,
		&limit.AmountAdminFee,
		&limit.Total,
		&limit.RatePerMonth,
		&limit.PayPerMonth,
		&limit.PercentageRate,
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

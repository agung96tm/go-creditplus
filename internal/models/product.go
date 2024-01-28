package models

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	PartnerID   int     `json:"partner_id"`
	PartnerName string  `json:"partner_name"`
}

type ProductModelInterface interface {
	GetAll() ([]*Product, error)
	Get(id int) (*Product, error)
}

type ProductModel struct {
	DB *sql.DB
}

func (m ProductModel) GetAll() ([]*Product, error) {
	query := `
		SELECT products.id, products.name, products.price, products.description, partners.id, partners.name
		FROM products
		INNER JOIN partners ON products.partner_id = partners.id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var products []*Product
	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Description,
			&product.PartnerID,
			&product.PartnerName,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (m ProductModel) Get(id int) (*Product, error) {
	query := `
		SELECT products.id, products.name, products.price, products.description, partners.id, partners.name
		FROM products
		INNER JOIN partners ON products.partner_id = partners.id
		WHERE products.id = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var product Product
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Description,
		&product.PartnerID,
		&product.PartnerName,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoDataFound
		default:
			return nil, err
		}
	}

	return &product, err
}

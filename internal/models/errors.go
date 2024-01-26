package models

import "errors"

var (
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrNoDataFound        = errors.New("models: data not found")
)

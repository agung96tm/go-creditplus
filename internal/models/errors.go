package models

import "errors"

var (
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrNoDataFound        = errors.New("models: data not found")
)

var (
	ErrLimitInsufficient = errors.New("models: Limit insufficient")
	ErrNoMatched         = errors.New("models: No Matched")
)

package domain

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrInternal        = errors.New("internal error")
	ErrRequired        = errors.New("required parameter is omitted")
	ErrInvalidPassword = errors.New("invalid password")
)

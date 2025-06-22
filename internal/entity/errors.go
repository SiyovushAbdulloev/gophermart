package entity

import "errors"

var (
	ErrNotFoundEntity         = errors.New("not found")
	ErrPasswordNotMatchEntity = errors.New("password not match")
)

package entity

import "errors"

var (
	NotFoundErrEntity         = errors.New("not found")
	PasswordNotMatchErrEntity = errors.New("password not match")
)

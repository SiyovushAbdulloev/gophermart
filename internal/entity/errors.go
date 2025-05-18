package entity

import "errors"

var (
	NotFoundErr         = errors.New("not found")
	PasswordNotMatchErr = errors.New("password not match")
)

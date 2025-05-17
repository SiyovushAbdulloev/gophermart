package repository

import "github.com/SiyovushAbdulloev/gophermart/internal/entity/user"

type AuthRepository interface {
	Register(user user.User) (*user.User, error)
	GetUserByEmail(email string) (*user.User, error)
}

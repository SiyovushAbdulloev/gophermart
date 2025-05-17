package usecase

import "github.com/SiyovushAbdulloev/gophermart/internal/entity/user"

type AuthUsecase interface {
	Register(user user.User) (string, error)
	GetUserByEmail(email string) (*user.User, error)
}

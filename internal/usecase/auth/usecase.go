package auth

import (
	"fmt"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/internal/repository"
	"github.com/SiyovushAbdulloev/gophermart/pkg/utils/jwt"
	"time"
)

type AuthUsecase struct {
	repo      repository.AuthRepository
	jwtSecret string
	jwtExpire time.Duration
}

func New(repo repository.AuthRepository, jwtSecret string, jwtExpire time.Duration) *AuthUsecase {
	return &AuthUsecase{
		repo:      repo,
		jwtSecret: jwtSecret,
		jwtExpire: jwtExpire,
	}
}

func (au *AuthUsecase) Register(user user.User) (string, error) {
	u, err := au.repo.Register(user)
	if err != nil {
		return "", err
	}

	token, err := jwt.JWTString(u.Id, au.jwtSecret, au.jwtExpire)
	if err != nil {
		fmt.Printf("Error creating token: %s\n", err.Error())
		return "", err
	}

	return token, nil
}

func (au *AuthUsecase) GetUserByEmail(email string) (*user.User, error) {
	return au.repo.GetUserByEmail(email)
}

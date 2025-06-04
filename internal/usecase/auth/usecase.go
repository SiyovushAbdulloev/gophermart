package auth

import (
	"fmt"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/internal/repository"
	"github.com/SiyovushAbdulloev/gophermart/pkg/utils/hash"
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

	token, err := jwt.JWTString(u.ID, au.jwtSecret, au.jwtExpire)
	if err != nil {
		fmt.Printf("Error creating token: %s\n", err.Error())
		return "", err
	}

	return token, nil
}

func (au *AuthUsecase) GetUserByEmail(email string) (*user.User, error) {
	return au.repo.GetUserByEmail(email)
}

func (au *AuthUsecase) Login(user *user.User) (string, error) {
	u, err := au.repo.GetUserByEmail(user.Email)
	if err != nil {
		return "", entity.NotFoundErrEntity
	}

	if !hash.CheckPassword(user.Password, u.Password) {
		return "", entity.PasswordNotMatchErrEntity
	}

	token, err := jwt.JWTString(u.ID, au.jwtSecret, au.jwtExpire)
	if err != nil {
		return "", err
	}

	return token, nil
}

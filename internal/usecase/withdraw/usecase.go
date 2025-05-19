package auth

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/withdraw"
	"github.com/SiyovushAbdulloev/gophermart/internal/repository"
)

type WithDrawUsecase struct {
	repo repository.WithDrawRepository
}

func New(repo repository.WithDrawRepository) *WithDrawUsecase {
	return &WithDrawUsecase{
		repo: repo,
	}
}

func (ou *WithDrawUsecase) List(userId int) ([]withdraw.WithDraw, error) {
	return ou.repo.List(userId)
}

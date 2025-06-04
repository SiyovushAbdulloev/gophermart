package auth

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
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

func (ou *WithDrawUsecase) List(userID int) ([]withdraw.WithDraw, error) {
	return ou.repo.List(userID)
}

func (ou *WithDrawUsecase) Store(w withdraw.WithDraw, u user.User) (*withdraw.WithDraw, error) {
	return ou.repo.Store(w, u)
}

func (ou *WithDrawUsecase) Sum(id int) (float64, error) {
	return ou.repo.Sum(id)
}

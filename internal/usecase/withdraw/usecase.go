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

func (ou *WithDrawUsecase) List(userId int) ([]withdraw.WithDraw, error) {
	return ou.repo.List(userId)
}

func (ou *WithDrawUsecase) Store(w withdraw.WithDraw, u user.User) (*withdraw.WithDraw, error) {
	return ou.repo.Store(w, u)
}

func (ou *WithDrawUsecase) Sum(id int) (int, error) {
	return ou.repo.Sum(id)
}

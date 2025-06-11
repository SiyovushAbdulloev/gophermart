package auth

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/repository"
)

type BalanceUsecase struct {
	repo repository.BalanceRepository
}

func New(repo repository.BalanceRepository) *BalanceUsecase {
	return &BalanceUsecase{
		repo: repo,
	}
}

func (bu *BalanceUsecase) GetAmount(id int) (float64, error) {
	return bu.repo.GetAmount(id)
}

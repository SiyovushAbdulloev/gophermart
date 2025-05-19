package auth

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/order"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/internal/repository"
)

type OrderUsecase struct {
	repo repository.OrderRepository
}

func New(repo repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{
		repo: repo,
	}
}

func (ou *OrderUsecase) Store(id int, u user.User) (*order.Order, error) {
	return ou.repo.Store(id, u)
}

func (ou *OrderUsecase) GetOrderById(id int) (*order.Order, error) {
	return ou.repo.GetOrderById(id)
}

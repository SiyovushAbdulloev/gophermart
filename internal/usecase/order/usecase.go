package order

import (
	"context"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/order"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/internal/repository"
)

type OrderUsecase struct {
	repo   repository.OrderRepository
	worker *Worker
}

func New(repo repository.OrderRepository, worker *Worker) *OrderUsecase {
	return &OrderUsecase{
		repo:   repo,
		worker: worker,
	}
}

func (ou *OrderUsecase) Store(id int, u user.User) (*order.Order, error) {
	o, err := ou.repo.Store(id, u)
	if err != nil {
		return nil, err
	}

	go ou.worker.Process(context.Background(), o)

	return o, nil
}

func (ou *OrderUsecase) GetOrderById(id int) (*order.Order, error) {
	return ou.repo.GetOrderById(id)
}

func (ou *OrderUsecase) List(userId int) ([]order.Order, error) {
	return ou.repo.List(userId)
}

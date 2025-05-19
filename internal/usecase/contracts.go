package usecase

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/order"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
)

type AuthUsecase interface {
	Register(user user.User) (string, error)
	GetUserByEmail(email string) (*user.User, error)
	Login(user *user.User) (string, error)
}

type OrderUsecase interface {
	Store(id int, u user.User) (*order.Order, error)
	GetOrderById(id int) (*order.Order, error)
	List(userId int) ([]order.Order, error)
}

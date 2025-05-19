package repository

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/order"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
)

type AuthRepository interface {
	Register(user user.User) (*user.User, error)
	GetUserByEmail(email string) (*user.User, error)
	GetUserById(id int) (*user.User, error)
}

type OrderRepository interface {
	Store(id int, u user.User) (*order.Order, error)
	GetOrderById(id int) (*order.Order, error)
}

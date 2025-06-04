package repository

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/order"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/withdraw"
)

type AuthRepository interface {
	Register(user user.User) (*user.User, error)
	GetUserByEmail(email string) (*user.User, error)
	GetUserByID(id int) (*user.User, error)
}

type OrderRepository interface {
	Store(id int, u user.User) (*order.Order, error)
	GetOrderByID(id int) (*order.Order, error)
	List(userID int) ([]order.Order, error)
	UpdateStatus(orderID int, status string) error
}

type WithDrawRepository interface {
	List(userId int) ([]withdraw.WithDraw, error)
	Store(w withdraw.WithDraw, u user.User) (*withdraw.WithDraw, error)
	Sum(id int) (int, error)
}

type BalanceRepository interface {
	GetAmount(id int) (int, error)
	AddPoints(userID int, amount int) error
}

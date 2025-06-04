package usecase

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/order"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/withdraw"
)

type AuthUsecase interface {
	Register(user user.User) (string, error)
	GetUserByEmail(email string) (*user.User, error)
	Login(user *user.User) (string, error)
}

type OrderUsecase interface {
	Store(id int, u user.User) (*order.Order, error)
	GetOrderByID(id int) (*order.Order, error)
	List(userID int) ([]order.Order, error)
}

type WithdrawUsecase interface {
	List(userID int) ([]withdraw.WithDraw, error)
	Store(w withdraw.WithDraw, u user.User) (*withdraw.WithDraw, error)
	Sum(id int) (int, error)
}

type BalanceUsecase interface {
	GetAmount(id int) (int, error)
}

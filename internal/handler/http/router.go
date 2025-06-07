package http

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/handler/http/auth"
	"github.com/SiyovushAbdulloev/gophermart/internal/handler/http/order"
	WithdrawHandler "github.com/SiyovushAbdulloev/gophermart/internal/handler/http/withdraw"
	"github.com/SiyovushAbdulloev/gophermart/internal/handler/middleware"
	"github.com/SiyovushAbdulloev/gophermart/internal/repository"
	"github.com/gin-gonic/gin"
)

func DefineAuthRoutes(app *gin.Engine, handler *auth.AuthHandler) {
	group := app.Group("/")
	group.Use(middleware.Guest())

	group.POST("/register", handler.Register)
	group.POST("/login", handler.Login)
}

func DefineOrderRoutes(app *gin.Engine, handler *order.OrderHandler, secret string, repository repository.AuthRepository) {
	group := app.Group("/orders")
	group.Use(middleware.Authenticate(secret, repository))

	group.POST("/", handler.Store)
	group.GET("/", handler.List)
}

func DefineWithdrawRoutes(app *gin.Engine, handler *WithdrawHandler.WithdrawHandler, secret string, repository repository.AuthRepository) {
	group := app.Group("/api")
	group.Use(middleware.Authenticate(secret, repository))

	group.GET("/user/withdrawals", handler.List)
	group.GET("/user/balance", handler.Balance)
	group.POST("/user/balance/withdraw", handler.Store)
}

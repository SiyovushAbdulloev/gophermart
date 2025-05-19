package http

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/handler/http/auth"
	"github.com/SiyovushAbdulloev/gophermart/internal/handler/http/order"
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

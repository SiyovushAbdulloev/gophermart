package http

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/handler/http/auth"
	"github.com/SiyovushAbdulloev/gophermart/internal/handler/middleware"
	"github.com/gin-gonic/gin"
)

func DefineAuthRoutes(app *gin.Engine, handler *auth.AuthHandler) {
	group := app.Group("/")
	group.Use(middleware.Guest())

	group.POST("/register", handler.Register)
	group.POST("/login", handler.Login)
}

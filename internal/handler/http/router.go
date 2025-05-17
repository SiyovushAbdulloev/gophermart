package http

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/handler/http/auth"
	"github.com/gin-gonic/gin"
)

func DefineAuthRoutes(app *gin.Engine, handler *auth.AuthHandler) {
	app.POST("/register", handler.Register)
}

package auth

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"io"
	"net/http"
)

type AuthHandler struct {
	uc usecase.AuthUsecase
}

func New(uc usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var u user.User

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = easyjson.Unmarshal(body, &u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existedUser, err := h.uc.GetUserByEmail(u.Email)
	if existedUser != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		return
	}

	token, err := h.uc.Register(u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var u user.User

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = easyjson.Unmarshal(body, &u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.uc.Login(&u)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

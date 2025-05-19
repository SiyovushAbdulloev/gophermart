package order

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WithdrawHandler struct {
	uc usecase.WithdrawUsecase
}

func New(uc usecase.WithdrawUsecase) *WithdrawHandler {
	return &WithdrawHandler{uc}
}

func (h *WithdrawHandler) List(ctx *gin.Context) {
	u, _ := ctx.Get("user")
	userData := u.(*user.User)

	withdraws, err := h.uc.List(userData.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"withdraws": withdraws,
	})
}

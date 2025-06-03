package withdraw

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/withdraw"
	"github.com/SiyovushAbdulloev/gophermart/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"io"
	"net/http"
)

type WithdrawHandler struct {
	uc        usecase.WithdrawUsecase
	balanceUC usecase.BalanceUsecase
}

func New(uc usecase.WithdrawUsecase, balanceUC usecase.BalanceUsecase) *WithdrawHandler {
	return &WithdrawHandler{
		uc,
		balanceUC,
	}
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

func (h *WithdrawHandler) Balance(ctx *gin.Context) {
	u, _ := ctx.Get("user")
	userData := u.(*user.User)

	balance, err := h.balanceUC.GetAmount(userData.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	withdrawn, err := h.uc.Sum(userData.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"current":   balance,
		"withdrawn": withdrawn,
	})
}

func (h *WithdrawHandler) Store(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var w withdraw.WithDraw
	err = easyjson.Unmarshal(body, &w)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, _ := ctx.Get("user")
	userData := u.(*user.User)

	balance, err := h.balanceUC.GetAmount(userData.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if int64(balance) < w.Sum {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	_, err = h.uc.Store(w, *userData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Withdraaw successfully stored",
	})
}

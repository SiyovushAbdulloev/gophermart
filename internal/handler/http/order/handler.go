package order

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/internal/usecase"
	"github.com/SiyovushAbdulloev/gophermart/pkg/utils/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"time"
)

type OrderHandler struct {
	uc usecase.OrderUsecase
}

func New(uc usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{uc}
}

func (h *OrderHandler) Store(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderID := string(body)

	if !utils.IsValidLuhn(orderID) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "wrong order"})
		return
	}

	oID, _ := strconv.Atoi(orderID)
	u, _ := ctx.Get("user")
	userData := u.(*user.User)

	previousOrder, _ := h.uc.GetOrderByID(oID)
	if previousOrder != nil {
		if previousOrder.UserID == userData.ID {
			ctx.JSON(http.StatusOK, gin.H{"message": "order already exists"})
		} else {
			ctx.JSON(http.StatusConflict, gin.H{"error": "order already exists by another user"})
		}
		return
	}

	_, err = h.uc.Store(oID, *userData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Order successfully stored",
	})
}

func (h *OrderHandler) List(ctx *gin.Context) {
	u, _ := ctx.Get("user")
	userData := u.(*user.User)

	orders, err := h.uc.List(userData.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type orderResponse struct {
		Number     string    `json:"number"`
		Status     string    `json:"status"`
		Accrual    float64   `json:"accrual,omitempty"`
		UploadedAt time.Time `json:"uploaded_at"`
	}

	ordersResp := make([]orderResponse, 0, len(orders))
	for _, o := range orders {
		resp := orderResponse{
			Number:     strconv.Itoa(o.ID),
			Status:     o.Status,
			UploadedAt: o.UpdatedAt,
		}
		if o.Status == "PROCESSED" {
			resp.Accrual = o.Points
		}
		ordersResp = append(ordersResp, resp)
	}

	ctx.JSON(http.StatusOK, ordersResp)
}

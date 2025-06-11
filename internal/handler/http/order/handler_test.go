package order

import (
	"bytes"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/order"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockOrderUsecase mocks the OrderUsecase interface
type mockOrderUsecase struct {
	mock.Mock
}

func (m *mockOrderUsecase) Store(orderID int, u user.User) (*order.Order, error) {
	args := m.Called(orderID, u)
	return args.Get(0).(*order.Order), args.Error(1)
}

func (m *mockOrderUsecase) GetOrderByID(orderID int) (*order.Order, error) {
	args := m.Called(orderID)
	return args.Get(0).(*order.Order), args.Error(1)
}

func (m *mockOrderUsecase) List(userID int) ([]order.Order, error) {
	args := m.Called(userID)
	return args.Get(0).([]order.Order), args.Error(1)
}

func TestOrderHandler_Store_Success(t *testing.T) {
	mockUC := new(mockOrderUsecase)
	handler := New(mockUC)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("user", &user.User{ID: 1})
		ctx.Next()
	})
	router.POST("/orders", handler.Store)

	mockUC.On("GetOrderById", 4242424242).Return((*order.Order)(nil), nil)
	mockUC.On("Store", 4242424242, user.User{ID: 1}).Return(&order.Order{ID: 1, Points: 4242424242, UserID: 1}, nil)

	body := []byte("4242424242")
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)
	mockUC.AssertExpectations(t)
}

func TestOrderHandler_Store_Conflict(t *testing.T) {
	mockUC := new(mockOrderUsecase)
	handler := New(mockUC)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("user", &user.User{ID: 1})
		ctx.Next()
	})
	router.POST("/orders", handler.Store)

	// Luhn-valid number
	mockUC.On("GetOrderById", 79927398713).Return(&order.Order{Points: 79927398713, UserID: 2}, nil)

	body := []byte("79927398713")
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	mockUC.AssertExpectations(t)
}

func TestOrderHandler_List_Success(t *testing.T) {
	mockUC := new(mockOrderUsecase)
	handler := New(mockUC)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("user", &user.User{ID: 1})
		ctx.Next()
	})
	router.GET("/orders", handler.List)

	mockUC.On("List", 1).Return([]order.Order{{ID: 1, Points: 123}}, nil)

	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUC.AssertExpectations(t)
}

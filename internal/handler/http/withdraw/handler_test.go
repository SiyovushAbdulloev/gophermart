package withdraw

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/withdraw"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// mockWithdrawUsecase mocks WithdrawUsecase
type mockWithdrawUsecase struct {
	mock.Mock
}

func (m *mockWithdrawUsecase) List(userID int) ([]withdraw.WithDraw, error) {
	args := m.Called(userID)
	return args.Get(0).([]withdraw.WithDraw), args.Error(1)
}

func (m *mockWithdrawUsecase) Store(w withdraw.WithDraw, u user.User) (*withdraw.WithDraw, error) {
	args := m.Called(w, u)
	return args.Get(0).(*withdraw.WithDraw), args.Error(1)
}

func (m *mockWithdrawUsecase) Sum(userID int) (float64, error) {
	args := m.Called(userID)
	return float64(args.Int(0)), args.Error(1)
}

// mockBalanceUsecase mocks BalanceUsecase
type mockBalanceUsecase struct {
	mock.Mock
}

func (m *mockBalanceUsecase) GetAmount(userID int) (float64, error) {
	args := m.Called(userID)
	return float64(args.Int(0)), args.Error(1)
}

func TestWithdrawHandler_List(t *testing.T) {
	uc := new(mockWithdrawUsecase)
	balanceUC := new(mockBalanceUsecase)
	handler := New(uc, balanceUC)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("user", &user.User{ID: 1})
		c.Next()
	})
	router.GET("/withdrawals", handler.List)

	uc.On("List", 1).Return([]withdraw.WithDraw{}, nil)

	req := httptest.NewRequest(http.MethodGet, "/withdrawals", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	uc.AssertExpectations(t)
}

func TestWithdrawHandler_Balance(t *testing.T) {
	uc := new(mockWithdrawUsecase)
	balanceUC := new(mockBalanceUsecase)
	handler := New(uc, balanceUC)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("user", &user.User{ID: 1})
		c.Next()
	})
	router.GET("/balance", handler.Balance)

	balanceUC.On("GetAmount", 1).Return(1000, nil)
	uc.On("Sum", 1).Return(250, nil)

	req := httptest.NewRequest(http.MethodGet, "/balance", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	balanceUC.AssertExpectations(t)
	uc.AssertExpectations(t)
}

func TestWithdrawHandler_Store(t *testing.T) {
	u := user.User{ID: 1}
	expectedWithdraw := withdraw.WithDraw{
		Order: "79927398713",
		Sum:   150,
	}

	uc := new(mockWithdrawUsecase)
	balanceUC := new(mockBalanceUsecase)

	// mock баланс
	balanceUC.On("GetAmount", u.ID).Return(200, nil)

	// mock Store с аргументами, на которые мы реагируем с помощью функции
	uc.On("Store", mock.MatchedBy(func(w withdraw.WithDraw) bool {
		return w.Order == expectedWithdraw.Order && w.Sum == expectedWithdraw.Sum
	}), u).Return(&expectedWithdraw, nil)

	handler := New(uc, balanceUC)

	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Set("user", &u)
		c.Next()
	})
	router.POST("/withdraw", handler.Store)

	body := `{"order": 79927398713, "sum": 150}`
	req := httptest.NewRequest(http.MethodPost, "/withdraw", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	uc.AssertExpectations(t)
	balanceUC.AssertExpectations(t)
}

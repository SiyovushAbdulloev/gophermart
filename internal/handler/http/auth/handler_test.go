package auth

import (
	"bytes"
	"errors"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockAuthUsecase mocks the AuthUsecase interface
type mockAuthUsecase struct {
	mock.Mock
}

func (m *mockAuthUsecase) Register(u user.User) (string, error) {
	args := m.Called(u)
	return args.String(0), args.Error(1)
}

func (m *mockAuthUsecase) Login(u *user.User) (string, error) {
	args := m.Called(u)
	return args.String(0), args.Error(1)
}

func (m *mockAuthUsecase) GetUserByEmail(email string) (*user.User, error) {
	args := m.Called(email)
	return args.Get(0).(*user.User), args.Error(1)
}

func TestAuthHandler_Register_Success(t *testing.T) {
	mockUC := new(mockAuthUsecase)
	handler := New(mockUC)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/register", handler.Register)

	mockUC.On("GetUserByEmail", "test@example.com").Return((*user.User)(nil), nil)
	mockUC.On("Register", mock.AnythingOfType("user.User")).Return("test-token", nil)

	body := `{"email":"test@example.com","password":"123456"}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUC.AssertExpectations(t)
}

func TestAuthHandler_Register_UserExists(t *testing.T) {
	mockUC := new(mockAuthUsecase)
	handler := New(mockUC)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/register", handler.Register)

	mockUC.On("GetUserByEmail", "test@example.com").Return(&user.User{Email: "test@example.com"}, nil)

	body := `{"email":"test@example.com","password":"123456"}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	mockUC.AssertExpectations(t)
}

func TestAuthHandler_Login_Success(t *testing.T) {
	mockUC := new(mockAuthUsecase)
	handler := New(mockUC)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", handler.Login)

	mockUC.On("Login", mock.AnythingOfType("*user.User")).Return("token-123", nil)

	body := `{"email":"test@example.com","password":"123456"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUC.AssertExpectations(t)
}

func TestAuthHandler_Login_Unauthorized(t *testing.T) {
	mockUC := new(mockAuthUsecase)
	handler := New(mockUC)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", handler.Login)

	mockUC.On("Login", mock.AnythingOfType("*user.User")).Return("", errors.New("unauthorized"))

	body := `{"email":"test@example.com","password":"wrongpass"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockUC.AssertExpectations(t)
}

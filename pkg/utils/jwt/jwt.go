package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func JWTString(id int, secret string, expiry time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiry).Unix(),
		},
		UserID: id,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckJWT(tokenString string, secret string) int {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("token parse error:", err)
		return -1
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return -1
	}

	return claims.UserID
}

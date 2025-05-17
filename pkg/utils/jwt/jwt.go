package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	*jwt.StandardClaims
	UserId int
}

func JWTString(id int, secret string, expiry time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiry).Unix(),
		},
		UserId: id,
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
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return -1
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return -1
	}

	return claims.UserId
}

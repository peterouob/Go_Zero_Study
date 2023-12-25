package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtPayload struct {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}

type CustomClaims struct {
	JwtPayload
	jwt.RegisteredClaims
}

func GetToken(user JwtPayload, accessSecret string, expires int64) (string, error) {
	claims := CustomClaims{user, jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour + time.Duration(expires)))}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessSecret))
}

func ParseToken(tokenStr string, accessSecret string, expires int64) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

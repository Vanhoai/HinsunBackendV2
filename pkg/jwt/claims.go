package jwt

import (
	jwtV4 "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	AccountID string `json:"accountId"`
	Email     string `json:"email"`
	JTI       string `json:"jti"`
	jwtV4.RegisteredClaims
}

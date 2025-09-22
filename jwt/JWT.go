package jwtUtil

import (
	jwtPkg "github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwtPkg.RegisteredClaims
}

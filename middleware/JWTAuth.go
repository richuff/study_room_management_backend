package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	jwtUtil "study_room_management_backend/jwt"
)

// JWTAuth 中间件：校验 Authorization Bearer token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 取 Header
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		fmt.Printf(tokenStr)
		// 2. 解析 & 校验
		claims := &jwtUtil.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtUtil.JwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 4. 把解析后的数据写回上下文，后续处理器可直接用
		c.Set("user_id", claims.Username)
		c.Set("claims", claims)

		c.Next()
	}
}

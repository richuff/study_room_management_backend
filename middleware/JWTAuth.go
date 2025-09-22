package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	jwtUtil "study_room_management_backend/jwt"
)

// JWTAuth 中间件：校验 Authorization Bearer token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 取 Header
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// 2. 剥离 Bearer
		parts := strings.SplitN(auth, " ", 2)
		if !(len(parts) == 2 && strings.EqualFold(parts[0], "Bearer")) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			c.Abort()
			return
		}
		tokenStr := parts[1]

		// 3. 解析 & 校验
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

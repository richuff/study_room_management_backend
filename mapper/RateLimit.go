package mapper

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ipMap = make(map[string]*ipLimit)
	mu    sync.Mutex
)

type ipLimit struct {
	lastTime time.Time
	count    int
}

const maxCount = 2

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP() // 获取用户真实IP

		mu.Lock()
		defer mu.Unlock()

		now := time.Now()
		// 拿到这个IP的记录
		limit, exists := ipMap[ip]

		if !exists {
			// 第一次访问
			ipMap[ip] = &ipLimit{
				lastTime: now,
				count:    1,
			}
			c.Next()
			return
		}

		// 判断是否在 1 秒内
		if now.Sub(limit.lastTime) < time.Second {
			limit.count++
			// 超过限制 → 直接拦截
			if limit.count > maxCount {
				c.JSON(http.StatusTooManyRequests, gin.H{
					"code": 429,
					"msg":  "操作过快，请稍后再试",
				})
				c.Abort()
				return
			}
		} else {
			// 超过1秒 → 重置计数
			limit.lastTime = now
			limit.count = 1
		}

		c.Next()
	}
}

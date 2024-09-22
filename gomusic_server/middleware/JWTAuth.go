package middleware

import (
	"github.com/gin-gonic/gin"
	"gomusic_server/common"
	"gomusic_server/utils"
	"net/http"
	"strings"
	"time"
)

// JWTAuthMiddleware 验证 JWT token 的中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, common.Error("请先登录！"))
			c.Abort()
			return
		}

		// 去掉前缀 "Bearer "
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token格式错误"})
			c.Abort()
			return
		}

		token := parts[1]

		// 检查是否在黑名单
		isBlacklisted, err := utils.IsTokenInBlacklist(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器错误"})
			c.Abort()
			return
		}

		if isBlacklisted {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token已失效"})
			c.Abort()
			return
		}

		// 解析 token 并验证其有效性
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的Token"})
			c.Abort()
			return
		}

		// 验证 token 是否过期
		if claims.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token已过期"})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文，方便后续使用
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// CORSMiddleware 前后端跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// 如果是预检请求，直接返回 204
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// 继续处理请求
		c.Next()
	}
}

// ReverseProxyMiddleware 将/music开头的请求转发到 MinIO
func ReverseProxyMiddleware(target *url.URL) gin.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(target)
	return func(c *gin.Context) {
		log.Print("进入minio静态资源映射....")
		if len(c.Request.URL.Path) > 6 && c.Request.URL.Path[:6] == "/music" {
			log.Printf("映射路径：" + c.Request.URL.Path)
			c.Request.URL.Scheme = target.Scheme
			c.Request.URL.Host = target.Host
			c.Request.Header.Set("X-Forwarded-Host", c.Request.Header.Get("Host"))
			proxy.ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return
		}
		c.Next()
	}
}

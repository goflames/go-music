package middleware

import (
	"gomusic_server/common"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// StaticFileMiddleware 处理前端请求静态资源时的映射关系
func StaticFileMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		switch {
		case strings.HasPrefix(path, "/img/"):
			// Handle /img static files
			log.Print("成功进入/img 静态资源映射....")
			filePath := "./assets" + path
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				c.JSON(http.StatusNotFound, common.Error("文件不存在"))
				c.Abort()
				return
			}
			c.File("./assets" + path)
			c.Abort()
			return

		case strings.HasPrefix(path, "/songSource/"):
			// Handle /songSource static files
			log.Print("成功进入/songSource/静态资源映射....")
			filepathParam := strings.TrimPrefix(path, "/songSource/")
			fullPath := filepath.Join("./assets/song", filepathParam)

			// Check if the file exists
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				c.JSON(http.StatusMethodNotAllowed, common.Error("该歌曲文件不存在！"))
				c.Abort()
				return
			}

			// Serve the file
			http.ServeFile(c.Writer, c.Request, fullPath)
			c.Abort()
			return

		case strings.HasPrefix(path, "/avatorImages/"):
			log.Print("成功进入/avatorImages/静态资源映射....")
			// Handle /avatorImages static files
			c.File("./asset/avatorImages" + path)
			c.Abort()
			return
		}

		// 路径不匹配时，继续执行其他程序
		c.Next()
	}
}

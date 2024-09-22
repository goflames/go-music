package router

import (
	"github.com/gin-gonic/gin"
	controller "gomusic_server/controllers"
	"gomusic_server/middleware"
	"net/url"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	//跨域配置
	router.Use(middleware.CORSMiddleware(),
		middleware.StaticFileMiddleware())

	// 应用反向代理中间件，将 /music/ 路径的请求转发到 MinIO
	minioURL, _ := url.Parse("http://localhost:9000")
	router.Use(middleware.ReverseProxyMiddleware(minioURL))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	bannerGroup := router.Group("/banner")
	{
		controller.BannerControllerRegister(bannerGroup)
	}

	songListGroup := router.Group("/songList")
	{
		controller.SongListControllerRegister(songListGroup)
	}

	singerGroup := router.Group("/singer")
	{
		controller.SingerControllerRegister(singerGroup)
	}

	songGroup := router.Group("/song")
	{
		controller.SongControllerRegister(songGroup)
	}

	listSongGroup := router.Group("/listSong")
	{
		controller.ListSongControllerRegister(listSongGroup)
	}

	rankListGroup := router.Group("/rankList")
	{
		controller.RankListControllerRegister(rankListGroup)
	}

	commentGroup := router.Group("/comment")
	{
		controller.CommentControllerRegister(commentGroup)
	}

	consumerGroup := router.Group("/user")
	{
		controller.ConsumerControllerRegister(consumerGroup)
	}

	collectionGroup := router.Group("/collection")
	{
		controller.CollectionControllerRegister(collectionGroup)
	}

	adminGroup := router.Group("/admin")
	{
		controller.AdminControllerRegister(adminGroup)
	}

	excelGroup := router.Group("/excle")
	{
		controller.ExcelControllerRegister(excelGroup)
	}

	downloadGroup := router.Group("/download")
	{
		controller.DownloadControllerRegister(downloadGroup)
	}

	userSupportGroup := router.Group("/userSupport")
	{
		controller.UsersupportControllerRegister(userSupportGroup)
	}
	return router
}

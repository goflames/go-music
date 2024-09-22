package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"gomusic_server/common"
	"gomusic_server/config"
	service "gomusic_server/services"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
)

type DownloadController struct {
	songService *service.SongService
}

func NewDownloadController(db *gorm.DB) *DownloadController {
	return &DownloadController{songService: service.NewSongService(db)}
}

func DownloadControllerRegister(router *gin.RouterGroup) {
	downloadController := NewDownloadController(config.DB)
	router.GET("/:fileName", downloadController.DownloadSong)
}

func (c *DownloadController) DownloadSong(ctx *gin.Context) {
	fileName := ctx.Param("fileName")
	objectName := "/songs/" + fileName
	object, err := config.MinioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Println("Failed to get object:", err)
		ctx.JSON(http.StatusOK, common.Error("Failed to download file"))
		return
	}
	defer object.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, object); err != nil {
		log.Println("Failed to read object:", err)
		ctx.JSON(http.StatusOK, common.Error("下载歌曲失败！"))
		return
	}
	fileBytes := buf.Bytes()

	// Set headers for file download
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Data(http.StatusOK, "application/octet-stream", fileBytes)
	ctx.JSON(http.StatusOK, "下载成功")
}

package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gomusic_server/common"
	"gomusic_server/config"
	"gomusic_server/models"
	service "gomusic_server/services"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

type SongController struct {
	songService *service.SongService
}

func NewSongController(db *gorm.DB) *SongController {
	return &SongController{service.NewSongService(db)}
}

func SongControllerRegister(router *gin.RouterGroup) {
	songController := NewSongController(config.DB)
	router.GET("/singer/detail", songController.GetSongsBySingerId)
	router.GET("/detail", songController.GetSongById)
	router.GET("", songController.GetAllSongs)
	router.GET("/singerName/detail", songController.GetSongBySingerName)
	router.POST("/img/update", songController.UpdateSongImg)
	router.POST("/add", songController.AddSong)
	router.DELETE("/delete", songController.DeleteById)
}

func (c *SongController) GetSongsBySingerId(ctx *gin.Context) {
	singerIdStr := ctx.Query("singerId")
	// 将 int64 转换为 int
	singerId, _ := strconv.Atoi(singerIdStr)

	// 调用查询
	songs, err := c.songService.GetSongsBySingerId(singerId)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("获取歌曲列表失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取歌曲列表成功", songs))

}

func (c *SongController) GetSongById(ctx *gin.Context) {
	idStr := ctx.Query("id")
	songId, _ := strconv.Atoi(idStr)
	song, err := c.songService.GetSongsById(songId)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("获取歌曲失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取歌曲成功", []models.Song{song}))
}

func (c *SongController) GetSongBySingerName(ctx *gin.Context) {
	name := ctx.Query("name")
	songs := c.songService.GetSongBySingerName(name)
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取歌曲成功", songs))
}

func (c *SongController) GetAllSongs(ctx *gin.Context) {
	songs := c.songService.GetAllSongs()
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取歌曲成功", songs))
}

func (c *SongController) UpdateSongImg(ctx *gin.Context) {
	idStr := ctx.Query("id")
	// 将 int64 转换为 int
	log.Print("idStr:" + idStr)
	id, _ := strconv.Atoi(idStr)
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "file is required"})
		return
	}

	prefix := "/songPic"
	objectName, err := service.UploadFileWithPrefix(file, bucketName, prefix)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("歌单照片更新失败"))
		return
	}
	response := c.songService.UpdateSongImg(id, "/"+bucketName+objectName)
	ctx.JSON(http.StatusOK, response)
}

func (c *SongController) AddSong(ctx *gin.Context) {
	name := ctx.PostForm("name")
	introduction := ctx.PostForm("introduction")
	singerIdStr := ctx.PostForm("singerId")
	singerId, _ := strconv.Atoi(singerIdStr)
	// Retrieve the MP3 file
	mpfile, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("MP3文件不能为空！"))
		return
	}

	// Retrieve the optional LRC file
	lrcfile, _ := ctx.FormFile("lrcfile") // ignore error as it's optional

	// Set static pic value
	pic := "/img/songPic/tubiao.jpg"

	// Upload MP3 file to MinIO
	prefix := "/songs"
	mpfileName, err := service.UploadFileWithPrefix(mpfile, bucketName, prefix)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("MP3文件上传失败"))
		return
	}

	// Prepare song object for database insertion
	song := models.Song{
		SingerID:     singerId,
		Name:         name,
		Introduction: introduction,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		Pic:          pic,
		URL:          "/" + bucketName + mpfileName,
	}

	// Handle LRC file if present
	if lrcfile != nil && song.Lyric == "[00:00:00]暂无歌词" {
		fileContent, err := lrcfile.Open()
		if err != nil {
			ctx.JSON(http.StatusOK, common.Error("LRC文件处理失败"))
			return
		}
		defer fileContent.Close()

		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(fileContent); err != nil {
			ctx.JSON(http.StatusOK, common.Error("LRC文件读取失败"))
			return
		}
		song.Lyric = buf.String()
	}

	// Insert song into the database
	response := c.songService.AddSong(song)

	// Return success response
	ctx.JSON(http.StatusOK, response)
}

func (c *SongController) DeleteById(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, _ := strconv.Atoi(idStr)
	objectName, isDeleted := c.songService.DeleteById(id)
	if isDeleted {
		err := service.RemoveFile(objectName, bucketName)
		if err != nil {
			ctx.JSON(http.StatusOK, common.Error("删除歌曲文件失败！"))
		}
	}

	ctx.JSON(http.StatusOK, common.Success("删除成功！"))
}

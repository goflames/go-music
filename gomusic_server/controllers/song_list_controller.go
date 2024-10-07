package controller

import (
	"github.com/gin-gonic/gin"
	"gomusic_server/common"
	"gomusic_server/config"
	"gomusic_server/dto"
	service "gomusic_server/services"
	"gomusic_server/utils"
	"gorm.io/gorm"
	"net/http"
)

const bucketName = "music"

type SongListController struct {
	songListService *service.SongListService
}

func NewSongListController(db *gorm.DB) *SongListController {
	return &SongListController{
		songListService: service.NewSongListDAOService(db)}
}

func SongListControllerRegister(router *gin.RouterGroup) {
	songListController := NewSongListController(config.DB)
	router.GET("", songListController.GetAllSongList)
	router.GET("/style/detail", songListController.GetSongListByStyle)
	router.GET("/likeTitle/detail", songListController.GetSongListLikeTitle)
	router.POST("/add", songListController.AddSongList)
	router.POST("/img/update", songListController.SongListUpdateImg)
	router.POST("/update", songListController.UpdateSongListInfo)
	router.GET("/delete", songListController.DeleteSongList)
}

func (c *SongListController) GetAllSongList(ctx *gin.Context) {
	songList, err := c.songListService.GetAllSongList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error("获取歌单列表失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取歌单列表成功", songList))
}

func (c *SongListController) GetSongListByStyle(ctx *gin.Context) {
	style := ctx.Query("style")
	songListByStyle, err := c.songListService.GetSongListByStyle(style)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error("获取歌单列表失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取歌单列表成功", songListByStyle))
}

// 根据歌单名模糊查询
func (c *SongListController) GetSongListLikeTitle(ctx *gin.Context) {
	title := ctx.Query("title")
	songList := c.songListService.GetSongListLikeTitle(title)
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取歌单列表成功", songList))
}

// 添加歌单
func (c *SongListController) AddSongList(ctx *gin.Context) {
	var request dto.SongListRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		return
	}
	if c.songListService.AddSongList(request) {
		ctx.JSON(http.StatusOK, common.Success("添加歌单成功"))
		return
	}
	ctx.JSON(http.StatusOK, common.Error("添加歌单失败"))
}

func (c *SongListController) SongListUpdateImg(ctx *gin.Context) {
	idStr := ctx.Query("id")
	// 将 int64 转换为 int8
	id := utils.TransferToInt8(idStr)
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"error": "file is required"})
		return
	}

	prefix := "/songListPic"
	objectName, err := service.UploadFileWithPrefix(file, bucketName, prefix)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("歌单照片更新失败"))
		return
	}
	response := c.songListService.UpdateSongListImg(id, "/"+bucketName+objectName)
	ctx.JSON(http.StatusOK, response)
}

func (c *SongListController) UpdateSongListInfo(ctx *gin.Context) {
	var request dto.SongListRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		return
	}
	response := c.songListService.UpdateSongListInfo(request)
	ctx.JSON(http.StatusOK, response)
}

func (c *SongListController) DeleteSongList(ctx *gin.Context) {
	idStr := ctx.Query("id")
	// 将 int64 转换为 int8
	id := utils.TransferToInt8(idStr)
	response := c.songListService.DeleteSongList(id)
	ctx.JSON(http.StatusOK, response)
}

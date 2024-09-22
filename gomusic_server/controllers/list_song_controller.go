package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gomusic_server/common"
	"gomusic_server/config"
	"gomusic_server/dto"
	"gomusic_server/models"
	service "gomusic_server/services"
	"gomusic_server/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ListSongController struct {
	listSongService *service.ListSongService
}

func NewListSongController(db *gorm.DB) *ListSongController {
	return &ListSongController{service.NewListSongService(db)}
}

func ListSongControllerRegister(router *gin.RouterGroup) {
	ListSongController := NewListSongController(config.DB)
	router.GET("/detail", ListSongController.GetSongsByListId)
	router.POST("/add", ListSongController.AddListSong)
	router.GET("/delete", ListSongController.DeleyeBySongId)
}

func (c *ListSongController) GetSongsByListId(ctx *gin.Context) {
	listIdStr := ctx.Query("songListId")
	listId := utils.TransferToInt8(listIdStr)
	listSongs, err := c.listSongService.GetSongsByListId(listId)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("获取歌单歌曲列表失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取歌单歌曲列表成功", listSongs))

}

func (c *ListSongController) AddListSong(ctx *gin.Context) {
	var request dto.ListSongRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		return
	}

	//songId, _ := StringToUint(request.SongID)
	songListId, _ := StringToUint(request.SongListID)

	var listSong models.ListSong
	listSong.SongID = request.SongID
	listSong.SongListID = songListId
	response := c.listSongService.AddListSong(listSong)
	ctx.JSON(http.StatusOK, response)
}

func (c *ListSongController) DeleyeBySongId(ctx *gin.Context) {
	songIdStr := ctx.Query("songListId")
	songId := utils.TransferToInt8(songIdStr)
	response := c.listSongService.DeleyeBySongId(songId)
	ctx.JSON(http.StatusOK, response)

}

// StringToUint 封装的将字符串转换为 uint 的方法
func StringToUint(s string) (uint, error) {
	// 将字符串解析为 uint64
	num, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, errors.New("转换错误: " + err.Error())
	}

	// 将 uint64 转换为 uint 类型
	return uint(num), nil
}

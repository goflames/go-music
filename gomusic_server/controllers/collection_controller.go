package controller

import (
	"github.com/gin-gonic/gin"
	"gomusic_server/common"
	"gomusic_server/config"
	"gomusic_server/dto"
	service "gomusic_server/services"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type CollectionController struct {
	collectionService *service.CollectionService
}

func NewCollectionController(db *gorm.DB) *CollectionController {
	return &CollectionController{
		service.NewCollectionService(db),
	}
}

func CollectionControllerRegister(router *gin.RouterGroup) {
	collectionController := NewCollectionController(config.DB)
	router.GET("/detail", collectionController.GetCollectionByUserId)
	router.POST("/status", collectionController.isCollected)
	router.POST("/add", collectionController.AddCollection)
	router.DELETE("/delete", collectionController.DeleteCollection)
}

func (c *CollectionController) GetCollectionByUserId(ctx *gin.Context) {
	userIdStr := ctx.Query("userId")
	// 将 int64 转换为 int
	userId, _ := strconv.Atoi(userIdStr)
	collections, err := c.collectionService.GetCollectionByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error("获取收藏歌单失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取收藏歌单成功！", collections))
}

func (c *CollectionController) isCollected(ctx *gin.Context) {
	var request dto.CollectRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		log.Print(err)
		return
	}
	collect := request.ToCollect()
	isCollect, err := c.collectionService.ExistSongID(collect)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error("获取收藏歌单失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取收藏歌单", isCollect))
}

func (c *CollectionController) AddCollection(ctx *gin.Context) {
	var request dto.CollectRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		log.Print(err)
		return
	}

	collect := request.ToCollect()
	response := c.collectionService.AddCollection(collect)
	ctx.JSON(http.StatusOK, response)
}

func (c *CollectionController) DeleteCollection(ctx *gin.Context) {
	userIdStr := ctx.Query("userId")
	songIdStr := ctx.Query("songId")
	userId, _ := strconv.Atoi(userIdStr)
	songId, _ := strconv.Atoi(songIdStr)
	response := c.collectionService.DeleteCollection(userId, songId)
	ctx.JSON(http.StatusOK, response)

}

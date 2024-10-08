package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomusic_server/common"
	"gomusic_server/config"
	"gomusic_server/services" // 导入服务层
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RankListController struct {
	rankListService *service.RankListService // 假设你有一个 RankListService 结构体处理业务逻辑
}

func NewRankListController(db *gorm.DB) *RankListController {
	return &RankListController{
		rankListService: service.NewRankListService(db)}
}

func RankListControllerRegister(router *gin.RouterGroup) {
	rankListController := NewRankListController(config.DB)
	router.GET("", rankListController.RankOfSongListId)
	router.GET("/user", rankListController.GetUserRank)
}

func (c *RankListController) RankOfSongListId(ctx *gin.Context) {
	songListIdStr := ctx.Query("songListId")
	songListId, _ := strconv.Atoi(songListIdStr)
	count, scoreSum, err := c.rankListService.RankOfSongListId(songListId)
	if count == 0 || scoreSum == 0 {
		ctx.JSON(http.StatusOK, common.Error("该歌单尚未获得评分"))
		return
	} else if err != nil {
		ctx.JSON(http.StatusOK, common.Error("获取评分失败"))
		return
	}
	// 计算平均值
	averageScore := float64(scoreSum) / float64(count)
	averageScoreFormatted := fmt.Sprintf("%.1f", averageScore)

	ctx.JSON(http.StatusOK, common.SuccessWithData("评分获取成功", averageScoreFormatted))
}

func (c *RankListController) GetUserRank(ctx *gin.Context) {
	consumerIdStr := ctx.Query("consumerId")
	songListIdStr := ctx.Query("songListId")

	consumerId, _ := strconv.Atoi(consumerIdStr)
	songListId, _ := strconv.Atoi(songListIdStr)
	rank, err := c.rankListService.GetUserRank(consumerId, songListId)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("该用户尚未评分"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("", rank))

}

package controller

import (
	"github.com/gin-gonic/gin"
	"gomusic_server/common"
	"gomusic_server/config"
	"gomusic_server/services"

	"gorm.io/gorm"
	"net/http"
)

type BannerController struct {
	bannerService *service.BannerService
}

func NewBannerController(db *gorm.DB) *BannerController {
	return &BannerController{
		bannerService: service.NewBannerService(db),
	}
}

func BannerControllerRegister(router *gin.RouterGroup) {
	bannerController := NewBannerController(config.DB)
	router.GET("/getAllBanner", bannerController.GetAllBanners)
}

func (c *BannerController) GetAllBanners(ctx *gin.Context) {
	banners, err := c.bannerService.GetAllBanners(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error("获取轮播图失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("成功获取轮播图", banners))
}

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
)

type AdminController struct {
	adminService *service.AdminService
}

func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{service.NewAdminService(db)}
}

func AdminControllerRegister(router *gin.RouterGroup) {
	adminController := NewAdminController(config.DB)
	router.POST("/login/status", adminController.AdminLoginStatus)
}

func (c *AdminController) AdminLoginStatus(ctx *gin.Context) {
	var adminRequest dto.AdminRequest
	if err := ctx.ShouldBindJSON(&adminRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		log.Print(err)
		return
	}

	status := c.adminService.AdminLoginStatus(adminRequest.Username, adminRequest.Password)
	if status {
		ctx.JSON(http.StatusOK, common.Success("管理员验证成功"))
		return
	}
	ctx.JSON(http.StatusOK, common.Error("管理员登陆失败"))

}

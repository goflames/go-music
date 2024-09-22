package controller

import (
	"github.com/gin-gonic/gin"
	"gomusic_server/common"
	"gomusic_server/config"
	"gomusic_server/dto"
	service "gomusic_server/services"
	"gorm.io/gorm"
	"net/http"
)

type UserSupportController struct {
	userSupportService *service.UserSupportService
}

// NewUserSupportController 创建新的 UserSupportController 实例
func NewUserSupportController(db *gorm.DB) *UserSupportController {
	return &UserSupportController{userSupportService: service.NewUserSupportService(db)}
}

func UsersupportControllerRegister(router *gin.RouterGroup) {
	userSupportController := NewUserSupportController(config.DB)
	router.POST("/test", userSupportController.IsUserSupportComment)
	router.POST("/insert", userSupportController.InsertCommentSupport)
	router.POST("/delete", userSupportController.DeleteCommentSupport)
}

// IsUserSupportComment 处理用户支持评论的请求
func (ctrl *UserSupportController) IsUserSupportComment(c *gin.Context) {
	var req dto.UserSupportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	response, err := ctrl.userSupportService.IsUserSupportComment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// InsertCommentSupport 处理添加记录请求
func (c *UserSupportController) InsertCommentSupport(ctx *gin.Context) {
	var request dto.UserSupportRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误: "+err.Error()))
		return
	}

	model := request.ToModel()

	response := c.userSupportService.InsertCommentSupport(model)
	ctx.JSON(http.StatusOK, response)
}

// DeleteCommentSupport 处理删除记录请求
func (c *UserSupportController) DeleteCommentSupport(ctx *gin.Context) {
	var request dto.UserSupportRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误: "+err.Error()))
		return
	}

	model := request.ToModel()

	response := c.userSupportService.DeleteCommentSupport(model)
	ctx.JSON(http.StatusOK, response)
}

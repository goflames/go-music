package controller

import (
	"github.com/gin-gonic/gin"
	"gomusic_server/common"
	"gomusic_server/config"
	"gomusic_server/dto"
	"gomusic_server/models"
	service "gomusic_server/services"
	"gomusic_server/utils"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type ConsumerController struct {
	consumerService *service.ConsumerService
}

func NewConsumerController(db *gorm.DB) *ConsumerController {
	return &ConsumerController{service.NewConsumerService(db)}
}

func ConsumerControllerRegister(router *gin.RouterGroup) {
	consumerController := NewConsumerController(config.DB)
	router.GET("", consumerController.GetAllUser)
	router.GET("/detail", consumerController.GetUserById)
	router.POST("/login/status", consumerController.LoginStatus)
	router.POST("/email/status", consumerController.EmailStatus)
	router.POST("/add", consumerController.Register)
	router.POST("/update", consumerController.UpdateConsumer)
	router.POST("/updatePassword", consumerController.UpdatePassword)
	router.POST("/avatar/update", consumerController.UpdateUserAvatar)
	router.GET("/delete", consumerController.DeleteById)
}

func (c *ConsumerController) GetUserById(ctx *gin.Context) {
	idStr := ctx.Query("id")
	// 将 int64 转换为 int8
	id := utils.TransferToInt8(idStr)
	consumer, err := c.consumerService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("获取评论用户失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取评论用户成功", []models.Consumer{consumer}))
}

// 用户账号密码登录
func (c *ConsumerController) LoginStatus(ctx *gin.Context) {
	var loginRequest dto.ConsumerRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		return
	}
	consumer, status, err := c.consumerService.LoginStatus(loginRequest)
	if consumer == nil {
		ctx.JSON(http.StatusOK, common.Error("用户不存在"))
		return
	} else if err != nil {
		ctx.JSON(http.StatusOK, common.Error("用户信息查询失败"))
		return
	}

	if !status {
		ctx.JSON(http.StatusOK, common.Error("用户名或密码错误"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("用户登录成功！", []*models.Consumer{consumer}))
}

// 用户邮箱登录
func (c *ConsumerController) EmailStatus(ctx *gin.Context) {
	var loginRequest dto.ConsumerRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		return
	}
	consumer, status, err := c.consumerService.EmailStatus(loginRequest)
	if consumer == nil {
		ctx.JSON(http.StatusOK, common.Error("用户不存在"))
		return
	} else if err != nil {
		ctx.JSON(http.StatusOK, common.Error("用户信息查询失败"))
		return
	}

	if !status {
		ctx.JSON(http.StatusOK, common.Error("用户名或密码错误"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("用户登录成功！", []*models.Consumer{consumer}))
}

func (c *ConsumerController) Register(ctx *gin.Context) {
	var request dto.ConsumerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		return
	}

	response, err := c.consumerService.AddUser(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *ConsumerController) UpdateConsumer(ctx *gin.Context) {
	var request dto.ConsumerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		return
	}

	response, err := c.consumerService.UpdateUser(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	ctx.JSON(http.StatusOK, response)

}

func (c *ConsumerController) UpdatePassword(ctx *gin.Context) {
	var request dto.ConsumerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		return
	}

	consumer, _ := c.consumerService.GetUserById(int8(request.Id))
	checkPassword := utils.CheckPassword(strings.TrimSpace(consumer.Password), strings.TrimSpace(request.OldPassword))

	if !checkPassword {
		ctx.JSON(http.StatusOK, common.Error("旧密码输入错误"))
		return
	}

	newPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, common.Error("密码加密失败"))
		return
	}
	consumer.Password = newPassword
	count, err := c.consumerService.UpdatePassword(consumer)
	if count < 1 {
		ctx.JSON(http.StatusUnauthorized, common.Error("更新密码失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.Success("更新成功"))
}

func (c *ConsumerController) DeleteById(ctx *gin.Context) {
	userIdStr := ctx.Query("id")
	userId := utils.TransferToInt8(userIdStr)

	if c.consumerService.DeleteById(userId) {
		ctx.JSON(http.StatusOK, common.Success("注销成功"))
		return
	}
	ctx.JSON(http.StatusOK, common.Error("更新失败"))
}

func (c *ConsumerController) GetAllUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.SuccessWithData("", c.consumerService.GetAllUser()))
}

func (c *ConsumerController) UpdateUserAvatar(ctx *gin.Context) {
	idStr := ctx.Query("id")
	// 将 int64 转换为 int8
	id := utils.TransferToInt8(idStr)
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"error": "file is required"})
		return
	}

	prefix := "/avatar"
	objectName, err := service.UploadFileWithPrefix(file, bucketName, prefix)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("用户头像更新失败"))
		return
	}
	response := c.consumerService.UpdateUserImg(id, "/"+bucketName+objectName)
	ctx.JSON(http.StatusOK, response)
}

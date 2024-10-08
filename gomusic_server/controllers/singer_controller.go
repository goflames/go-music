package controller

import (
	"github.com/gin-gonic/gin"
	"gomusic_server/common"
	"gomusic_server/config"
	"gomusic_server/dto"
	"gomusic_server/models"
	service "gomusic_server/services"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type SingerController struct {
	singerService *service.SingerService
}

func NewSingerController(db *gorm.DB) *SingerController {
	return &SingerController{
		singerService: service.NewSingerService(db),
	}
}

func SingerControllerRegister(router *gin.RouterGroup) {
	singerController := NewSingerController(config.DB)
	router.GET("", singerController.GetAllSingers)
	router.GET("/sex/detail", singerController.GetSingersByGerder)
	router.POST("/update", singerController.UpdateSingerInfo)
	router.POST("/add", singerController.AddSinger)
	router.POST("/avatar/update", singerController.UpdateSingerImg)
	router.DELETE("/delete", singerController.DeleteSingerById)
}

func (c *SingerController) GetAllSingers(ctx *gin.Context) {
	singers, err := c.singerService.GetAllSingers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error("获取歌手列表失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取歌手列表成功", singers))
}

func (c *SingerController) GetSingersByGerder(ctx *gin.Context) {
	genderStr := ctx.Query("sex")
	// 将 int64 转换为 int
	gender, _ := strconv.Atoi(genderStr)

	singers, err := c.singerService.GetSingersByGender(gender)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("获取歌手列表失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取歌手列表成功", singers))
}

func (c *SingerController) UpdateSingerInfo(ctx *gin.Context) {
	var request dto.SingerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		return
	}
	err := c.singerService.UpdateSingerInfo(request)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("更新信息失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.Success("更新信息成功"))
}

func (c *SingerController) UpdateSingerImg(ctx *gin.Context) {
	idStr := ctx.Query("id")
	// 将 int64 转换为 int
	id, _ := strconv.Atoi(idStr)
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"error": "file is required"})
		return
	}

	prefix := "/singerPic"
	objectName, err := service.UploadFileWithPrefix(file, bucketName, prefix)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error("歌单照片更新失败"))
		return
	}
	response := c.singerService.UpdateSingerImg(id, "/"+bucketName+objectName)
	ctx.JSON(http.StatusOK, response)
}

func (c *SingerController) DeleteSingerById(ctx *gin.Context) {
	idStr := ctx.Query("id")
	// 将 int64 转换为 int
	id, _ := strconv.Atoi(idStr)

	// todo：确认song表中是否还有该歌手的歌曲

	response := c.singerService.DeleteSingerById(id)
	ctx.JSON(http.StatusOK, response)
}

func (c *SingerController) AddSinger(ctx *gin.Context) {
	var request dto.SingerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		return
	}

	birth, _ := time.Parse("2006-01-02", request.Birth)
	singer := models.Singer{
		Name:         request.Name,
		Sex:          request.Sex,
		Pic:          "/img/avatorImages/user.jpg",
		Location:     request.Location,
		Introduction: request.Introduction,
		Birth:        birth,
	}
	response := c.singerService.AddSinger(singer)
	ctx.JSON(http.StatusOK, response)
}

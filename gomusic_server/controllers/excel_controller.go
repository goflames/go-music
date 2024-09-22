package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomusic_server/config"
	service "gomusic_server/services"
	"gomusic_server/utils"
	"io"
	"net/http"
	"os"
	"time"
)

type ExcelController struct {
}

func NewExcelController() *ExcelController {
	return &ExcelController{}
}

func ExcelControllerRegister(router *gin.RouterGroup) {
	excelController := NewExcelController()
	router.GET("", excelController.OutPutExcel)
}

func (c *ExcelController) OutPutExcel(ctx *gin.Context) {
	fileName := fmt.Sprintf("SongList%d.xlsx", time.Now().Unix())
	songListService := service.NewSongListDAOService(config.DB)
	allSongList, _ := songListService.GetAllSongList()

	// 调用通用的 createExcel 方法生成文件
	err := utils.CreateExcel(fileName, allSongList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create Excel file",
		})
		return
	}

	// 打开生成的 Excel 文件
	file, err := os.Open(fileName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to open Excel file",
		})
		return
	}
	defer file.Close()

	// 设置响应头，准备下载文件
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Length", fmt.Sprintf("%d", utils.GetFileSize(file)))

	// 将文件发送到客户端
	_, err = io.Copy(ctx.Writer, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send Excel file",
		})
		return
	}
}

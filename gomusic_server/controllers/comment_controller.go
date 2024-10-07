package controller

import (
	"github.com/gin-gonic/gin"
	"gomusic_server/common"
	"gomusic_server/config"
	"gomusic_server/dto"
	service "gomusic_server/services"
	"gomusic_server/utils"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController(db *gorm.DB) *CommentController {
	return &CommentController{commentService: service.NewCommentService(db)}
}

func CommentControllerRegister(router *gin.RouterGroup) {
	commentController := NewCommentController(config.DB)
	router.GET("/songList/detail", commentController.CommentOfSongListId)
	router.GET("/song/detail", commentController.CommentOfSongId)
	router.POST("/add", commentController.AddComment)
	router.POST("/like", commentController.CommentOfLike)
	router.GET("/delete", commentController.DeleteComment)
}

func (c *CommentController) CommentOfSongListId(ctx *gin.Context) {
	songListIdStr := ctx.Query("songListId")
	// 将 int64 转换为 int8
	songListId := utils.TransferToInt8(songListIdStr)
	comments, err := c.commentService.CommentOfSongListId(songListId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error("获取评论失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取评论成功", comments))
}

func (c *CommentController) CommentOfSongId(ctx *gin.Context) {
	songIdStr := ctx.Query("songId")
	// 将 int64 转换为 int8
	songId := utils.TransferToInt8(songIdStr)
	comments, err := c.commentService.CommentOfSongId(songId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error("获取评论失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.SuccessWithData("获取评论成功", comments))
}

// 新增评论
func (c *CommentController) AddComment(ctx *gin.Context) {
	var request dto.CommentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		log.Print(err)
		return
	}

	comment := request.ToComment()

	isAdd := c.commentService.AddComment(comment)

	if isAdd {
		ctx.JSON(http.StatusOK, common.Success("评论成功"))
		return
	}
	ctx.JSON(http.StatusBadRequest, common.Error("评论失败"))
}

// CommentController 处理点赞请求
func (c *CommentController) CommentOfLike(ctx *gin.Context) {
	var commentRequest dto.CommentRequest
	// 从请求中绑定 JSON 数据
	if err := ctx.ShouldBindJSON(&commentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error("请求参数错误"))
		return
	}

	// 调用 service 方法进行点赞更新
	response := c.commentService.UpdateCommentMsg(commentRequest)
	ctx.JSON(http.StatusOK, response)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	idStr := ctx.Query("id")
	// 将 int64 转换为 int8
	id := utils.TransferToInt8(idStr)
	response := c.commentService.DeleteComment(id)
	ctx.JSON(http.StatusOK, response)
}

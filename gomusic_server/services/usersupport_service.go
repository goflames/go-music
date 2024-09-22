package service

import (
	"gomusic_server/common"
	"gomusic_server/dto"
	"gomusic_server/models"
	"gorm.io/gorm"
)

type UserSupportService struct {
	db *gorm.DB
}

// NewUserSupportService 创建新的 UserSupportService 实例
func NewUserSupportService(db *gorm.DB) *UserSupportService {
	return &UserSupportService{db: db}
}

// IsUserSupportComment 检查用户是否点赞评论
func (s *UserSupportService) IsUserSupportComment(req dto.UserSupportRequest) (common.Response, error) {
	var count int64
	err := s.db.Model(&models.UserSupport{}).
		Where("comment_id = ? AND user_id = ?", req.CommentID, req.UserID).
		Count(&count).Error

	if err != nil {
		return common.Error("获取点赞数据失败"), err
	}

	if count > 0 {
		return common.SuccessWithData("您已取消点赞", true), nil
	}
	return common.SuccessWithData("点赞成功", false), nil
}

// InsertCommentSupport 添加记录
func (s *UserSupportService) InsertCommentSupport(userSupportRequest *models.UserSupport) common.Response {
	if err := s.db.Create(userSupportRequest).Error; err != nil {
		return common.Error("添加记录时发生异常: " + err.Error())
	}
	return common.Success("添加记录成功")
}

// DeleteCommentSupport 删除记录
func (s *UserSupportService) DeleteCommentSupport(userSupportRequest *models.UserSupport) common.Response {
	tx := s.db.Where("comment_id = ? AND user_id = ?", userSupportRequest.CommentID, userSupportRequest.UserID).Delete(&models.UserSupport{})
	if tx.RowsAffected < 1 || tx.Error != nil {
		return common.Error("删除记录失败！")
	}
	return common.Success("删除记录成功")
}

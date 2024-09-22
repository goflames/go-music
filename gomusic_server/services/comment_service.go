package service

import (
	"gomusic_server/common"
	"gomusic_server/dao"
	"gomusic_server/dto"
	"gomusic_server/models"
	"gorm.io/gorm"
)

type CommentService struct {
	commentDAO *dao.CommentDAO
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{commentDAO: dao.NewCommentDAO(db)}
}

func (s *CommentService) CommentOfSongListId(songListId int8) ([]models.Comment, error) {
	comments, err := s.commentDAO.CommentOfSongListId(songListId)
	return comments, err
}

func (s *CommentService) CommentOfSongId(songId int8) ([]models.Comment, error) {
	comments, err := s.commentDAO.CommentOfSongId(songId)
	return comments, err
}

func (s *CommentService) AddComment(comment models.Comment) bool {
	return s.commentDAO.AddComment(comment)
}

// UpdateCommentMsg 更新评论点赞信息
func (s *CommentService) UpdateCommentMsg(commentRequest dto.CommentRequest) common.Response {
	// 将 CommentRequest 转换为 Comment
	comment := commentRequest.ToComment()
	// 执行数据库更新操作
	if err := s.commentDAO.UpdateComment(&comment); err != nil {
		return common.Error("点赞失败")
	}
	return common.Success("点赞成功")
}

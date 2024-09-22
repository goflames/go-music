package dao

import (
	"gomusic_server/models"
	"gorm.io/gorm"
)

type CommentDAO struct {
	db *gorm.DB
}

func NewCommentDAO(db *gorm.DB) *CommentDAO {
	return &CommentDAO{db}
}

func (dao *CommentDAO) CommentOfSongListId(songListId int8) ([]models.Comment, error) {
	var comments []models.Comment
	err := dao.db.Where("song_list_id = ?", songListId).Find(&comments).Error
	return comments, err
}

func (dao *CommentDAO) CommentOfSongId(songId int8) ([]models.Comment, error) {
	var comments []models.Comment
	err := dao.db.Where("song_id = ?", songId).Find(&comments).Error
	return comments, err
}

func (dao *CommentDAO) AddComment(comment models.Comment) bool {
	tx := dao.db.Create(&comment)
	return tx.RowsAffected > 0
}

// UpdateComment 更新评论信息
func (dao *CommentDAO) UpdateComment(comment *models.Comment) error {
	return dao.db.Model(&models.Comment{}).Where("id = ?", comment.ID).Updates(comment).Error
}

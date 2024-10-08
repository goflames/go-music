package dao

import (
	"gomusic_server/common"
	"gomusic_server/models"
	"gorm.io/gorm"
	"log"
)

type CommentDAO struct {
	db *gorm.DB
}

func NewCommentDAO(db *gorm.DB) *CommentDAO {
	return &CommentDAO{db}
}

func (dao *CommentDAO) CommentOfSongListId(songListId int) ([]models.Comment, error) {
	var comments []models.Comment
	err := dao.db.Where("song_list_id = ?", songListId).Find(&comments).Error
	return comments, err
}

func (dao *CommentDAO) CommentOfSongId(songId int) ([]models.Comment, error) {
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

func (dao *CommentDAO) DeleteComment(id int) common.Response {
	// 通过主键删除
	tx := dao.db.Delete(&models.Comment{}, id)
	if tx.Error != nil || tx.RowsAffected < 1 {
		log.Print("删除评论失败....")
		return common.Error("删除歌单失败！请重试！")
	}
	return common.Success("删除歌单成功")
}

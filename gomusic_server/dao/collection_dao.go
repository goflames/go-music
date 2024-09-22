package dao

import (
	"gomusic_server/common"
	"gomusic_server/models"
	"gorm.io/gorm"
)

type CollectionDAO struct {
	db *gorm.DB
}

func NewCollectionDAO(db *gorm.DB) *CollectionDAO {
	return &CollectionDAO{db}
}

func (dao *CollectionDAO) GetCollectionByUserId(userId int8) ([]*models.Collection, error) {
	var collections []*models.Collection
	tx := dao.db.Where("user_id = ?", userId).Find(&collections)
	return collections, tx.Error
}

func (dao *CollectionDAO) ExistSongID(userId, songId int) (int, error) {
	var collections []*models.Collection
	tx := dao.db.Where("user_id = ? and song_id = ?", userId, songId).Find(&collections)
	return len(collections), tx.Error
}

func (dao *CollectionDAO) AddCollection(collect models.Collection) common.Response {
	tx := dao.db.Create(&collect)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return common.Error("收藏失败！")
	}
	return common.SuccessWithData("收藏成功！", true)
}

func (dao *CollectionDAO) DeleteCollection(userId, songId int) common.Response {
	tx := dao.db.Where("user_id = ? and song_id = ?", userId, songId).Delete(&models.Collection{})
	if tx.Error == nil && tx.RowsAffected > 0 {
		return common.SuccessWithData("取消收藏！", false)
	}
	return common.Error("取消收藏失败！")
}

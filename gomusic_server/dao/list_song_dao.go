package dao

import (
	"gomusic_server/common"
	"gomusic_server/models"
	"gorm.io/gorm"
	"strings"
)

type ListSongDAO struct {
	db *gorm.DB
}

func NewListSongDAO(db *gorm.DB) *ListSongDAO {
	return &ListSongDAO{db}
}

func (dao *ListSongDAO) GetSongsByListId(listId int) ([]models.ListSong, error) {
	var listSong []models.ListSong
	err := dao.db.Where("song_list_id = ?", listId).Find(&listSong).Error
	return listSong, err
}

func (dao *ListSongDAO) AddListSong(listSong models.ListSong) common.Response {
	result := dao.db.Create(&listSong)
	//检查是否有重复数据（通过唯一约束引发的错误）
	if result.Error != nil {
		// 检查是否是 MySQL 唯一约束冲突错误
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			return common.Error("歌单中已存在该歌曲")
		}
		// 处理其他错误
		return common.Error("加入歌单失败，请重试")
	}
	return common.Success("添加成功！")
}

func (dao *ListSongDAO) DeleyeBySongId(songId int) common.Response {
	result := dao.db.Delete(&models.ListSong{}, songId)
	if result.Error != nil && result.RowsAffected > 0 {
		return common.Error("删除失败！")
	}
	return common.Success("删除成功！")
}

package dao

import (
	"gomusic_server/common"
	"gomusic_server/models"
	"gorm.io/gorm"
	"log"
	"strings"
)

type SongDAO struct {
	db *gorm.DB
}

func NewSongDAO(db *gorm.DB) *SongDAO {
	return &SongDAO{db}
}

func (dao *SongDAO) GetSongsBySingerId(singerId int) ([]models.Song, error) {
	var songs []models.Song
	err := dao.db.Where("singer_id = ?", singerId).Find(&songs).Error
	return songs, err
}

func (dao *SongDAO) GetSongsById(songId int) (models.Song, error) {
	var song models.Song
	err := dao.db.First(&song, songId).Error
	return song, err
}

func (dao *SongDAO) GetSongBySingerName(name string) []models.Song {
	var songs []models.Song
	dao.db.Where("name != '' AND name like ?", "%"+name+"%").Find(&songs)
	return songs
}

func (dao *SongDAO) GetAllSongs() ([]models.Song, error) {
	var songs []models.Song
	tx := dao.db.Find(&songs)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return songs, nil
}

func (dao *SongDAO) UpdateSongImg(songId int, pic string) common.Response {
	tx := dao.db.Model(&models.Song{}).Where("id = ?", songId).Update("pic", pic)
	if tx.Error != nil || tx.RowsAffected < 1 {
		log.Print("歌曲图片更新失败....")
		return common.Error("数据库更新歌曲图片失败")
	}
	return common.SuccessWithData("更新成功", pic)
}

func (dao *SongDAO) AddSong(song models.Song) common.Response {
	tx := dao.db.Create(&song)
	if tx.Error != nil && tx.RowsAffected > 0 {
		return common.Error("插入数据失败")
	}
	return common.SuccessWithData("新增歌曲成功", song.URL)
}

func (dao *SongDAO) DeleteById(id int) (string, bool) {
	song, err := dao.GetSongsById(id)
	if err != nil {
		return "", false
	}
	tx := dao.db.Delete(&song)
	if tx.Error == nil && tx.RowsAffected < 1 {
		return "", false
	}
	// 找到第一个 '/' 的位置
	firstIndex := strings.Index(song.URL, "/")

	// 在第一个 '/' 之后的子字符串中，找到第二个 '/' 的位置
	secondIndex := strings.Index(song.URL[firstIndex+1:], "/") + firstIndex + 1

	// 从第二个 '/' 开始截取
	objectName := song.URL[secondIndex:]
	return objectName, true
}

func (dao *SongDAO) GetSongsByName(songName string) (models.Song, error) {
	var song models.Song
	tx := dao.db.Where("name = ?", songName).Find(&song)
	return song, tx.Error
}

package dao

import (
	"gomusic_server/common"
	"gomusic_server/dto"
	"gomusic_server/models"
	"gorm.io/gorm"
)

type SongListDAO struct {
	db *gorm.DB
}

func NewSongListDAO(db *gorm.DB) *SongListDAO {
	return &SongListDAO{db: db}
}

func (dao *SongListDAO) GetAllSongList() ([]models.SongList, error) {
	var songList []models.SongList
	if err := dao.db.Find(&songList).Error; err != nil {
		return nil, err
	}
	return songList, nil
}

// 根据查询风格歌单列表
func (dao *SongListDAO) GetSongListByStyle(style string) ([]models.SongList, error) {
	var songLists []models.SongList
	err := dao.db.Where("style LIKE ?", "%"+style+"%").Find(&songLists).Error
	return songLists, err
}

func (dao *SongListDAO) GetSongListLikeTitle(title string) []models.SongList {
	var songLists []models.SongList
	dao.db.Where("title != '' AND title like ?", "%"+title+"%").Find(&songLists)
	return songLists
}

func (dao *SongListDAO) AddSongList(request dto.SongListRequest) error {
	var songlist models.SongList
	songlist.Title = request.Title
	songlist.Pic = request.Pic
	songlist.Style = request.Style
	songlist.Introduction = request.Introduction
	err := dao.db.Create(&songlist).Error
	return err
}

func (dao *SongListDAO) UpdateSongListImg(id int8, img string) common.Response {
	// 通过主键更新某个字段
	tx := dao.db.Model(&models.SongList{}).Where("id = ?", id).Update("pic", img)
	if tx.Error != nil {
		return common.Error("更新歌单图片失败！")
	}
	return common.SuccessWithData("更新歌单图片成功！", img)
}

func (dao *SongListDAO) UpdateSongListInfo(request dto.SongListRequest) common.Response {
	var songList models.SongList
	dao.db.First(&songList, request.ID)
	songList.Title = request.Title
	songList.Style = request.Style
	songList.Introduction = request.Introduction
	tx := dao.db.Save(&songList).Omit("pic")
	if tx.Error != nil || tx.RowsAffected < 1 {
		return common.Error("更新歌单信息失败")
	}
	return common.Success("更新歌单信息成功！")
}

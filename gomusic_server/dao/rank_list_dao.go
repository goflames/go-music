package dao

import (
	"gomusic_server/models"
	"gorm.io/gorm"
)

type RankListDAO struct {
	Db *gorm.DB
}

func NewRankListDAO(db *gorm.DB) *RankListDAO {
	return &RankListDAO{Db: db}
}

func (dao *RankListDAO) RankOfSongListId(songListId int8) (int64, float64, error) {
	var count int64
	var totalScore float64

	// 获取记录总数
	err := dao.Db.Model(&models.RankList{}).
		Where("song_list_id = ?", songListId).
		Count(&count).Error
	if err != nil {
		return 0, 0, err
	}

	// 获取所有评分的总和
	err = dao.Db.Model(&models.RankList{}).
		Where("song_list_id = ?", songListId).
		Select("SUM(score)").
		Scan(&totalScore).Error
	if err != nil {
		return 0, 0, err
	}

	return count, totalScore, nil
}

func (dao *RankListDAO) GetUserRank(consumerId int8, songlistId int8) (uint32, error) {
	var rankList models.RankList
	err := dao.Db.Where("consumer_id = ? and song_list_id = ?", consumerId, songlistId).Find(&rankList).Error
	return rankList.Score, err
}

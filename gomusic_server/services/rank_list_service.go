package service

import (
	"gomusic_server/dao"
	"gorm.io/gorm"
)

type RankListService struct {
	rankListDAO *dao.RankListDAO
}

func NewRankListService(db *gorm.DB) *RankListService {
	return &RankListService{
		rankListDAO: dao.NewRankListDAO(db),
	}
}

func (s *RankListService) RankOfSongListId(songListId int) (int64, float64, error) {
	count, scoreSum, err := s.rankListDAO.RankOfSongListId(songListId)
	return count, scoreSum, err
}

func (s *RankListService) GetUserRank(consumerid int, songListId int) (uint32, error) {
	score, err := s.rankListDAO.GetUserRank(consumerid, songListId)
	return score, err
}

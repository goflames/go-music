package service

import (
	"gomusic_server/common"
	"gomusic_server/dao"
	"gomusic_server/dto"
	"gomusic_server/models"
	"gorm.io/gorm"
)

type SongListService struct {
	songListDAO *dao.SongListDAO
}

func NewSongListDAOService(db *gorm.DB) *SongListService {
	return &SongListService{dao.NewSongListDAO(db)}
}

func (s *SongListService) GetAllSongList() ([]models.SongList, error) {
	songList, err := s.songListDAO.GetAllSongList()
	if err != nil {
		return nil, err
	}
	return songList, nil
}

func (s *SongListService) GetSongListByStyle(style string) ([]models.SongList, error) {
	songList, err := s.songListDAO.GetSongListByStyle(style)
	return songList, err
}

func (s *SongListService) GetSongListLikeTitle(title string) []models.SongList {
	return s.songListDAO.GetSongListLikeTitle(title)
}

func (s *SongListService) AddSongList(request dto.SongListRequest) bool {
	err := s.songListDAO.AddSongList(request)
	if err != nil {
		return false
	}
	return true
}

func (s *SongListService) UpdateSongListImg(id int8, img string) common.Response {
	return s.songListDAO.UpdateSongListImg(id, img)
}

func (s *SongListService) UpdateSongListInfo(request dto.SongListRequest) common.Response {
	return s.songListDAO.UpdateSongListInfo(request)
}

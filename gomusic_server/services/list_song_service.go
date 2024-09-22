package service

import (
	"gomusic_server/common"
	"gomusic_server/dao"
	"gomusic_server/models"
	"gorm.io/gorm"
)

type ListSongService struct {
	listSongDao *dao.ListSongDAO
}

func NewListSongService(db *gorm.DB) *ListSongService {
	return &ListSongService{dao.NewListSongDAO(db)}
}

func (s *ListSongService) GetSongsByListId(listId int8) ([]models.ListSong, error) {
	listSongs, err := s.listSongDao.GetSongsByListId(listId)
	return listSongs, err
}

func (s *ListSongService) AddListSong(listSong models.ListSong) common.Response {
	return s.listSongDao.AddListSong(listSong)
}

func (s *ListSongService) DeleyeBySongId(songId int8) common.Response {
	return s.listSongDao.DeleyeBySongId(songId)
}

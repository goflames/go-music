package service

import (
	"gomusic_server/common"
	"gomusic_server/dao"
	"gomusic_server/models"
	"gorm.io/gorm"
)

type SongService struct {
	songDao *dao.SongDAO
}

func NewSongService(db *gorm.DB) *SongService {
	return &SongService{dao.NewSongDAO(db)}
}

func (s *SongService) GetSongsBySingerId(singerId int) ([]models.Song, error) {
	songs, err := s.songDao.GetSongsBySingerId(singerId)
	return songs, err
}

func (s *SongService) GetSongsById(songId int) (models.Song, error) {
	song, err := s.songDao.GetSongsById(songId)
	return song, err
}

func (s *SongService) GetSongBySingerName(name string) []models.Song {
	return s.songDao.GetSongBySingerName(name)
}

func (s *SongService) GetAllSongs() []models.Song {
	songs, err := s.songDao.GetAllSongs()
	if err != nil {
		return nil
	}
	return songs
}

func (s *SongService) UpdateSongImg(songId int, pic string) common.Response {
	return s.songDao.UpdateSongImg(songId, pic)
}

func (s *SongService) AddSong(song models.Song) common.Response {
	return s.songDao.AddSong(song)
}

func (s *SongService) DeleteById(id int) (string, bool) {
	return s.songDao.DeleteById(id)
}

func (s *SongService) GetSongsByName(songName string) (models.Song, error) {
	song, err := s.songDao.GetSongsByName(songName)
	return song, err
}

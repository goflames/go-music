package service

import (
	"gomusic_server/common"
	"gomusic_server/dao"
	"gomusic_server/dto"
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

func (s *SongService) UpdateSongInfo(request dto.SongRequest) error {
	return s.songDao.UpdateSongInfo(request)
}

func (s *SongService) UpdateSongLrc(lrcFile []byte, id int) (string, error) {
	song, err := s.songDao.GetSongByID(id)
	if err != nil {
		return "获取歌曲失败", err
	}

	if lrcFile != nil && song.Lyric != "[00:00:00]暂无歌词" {
		// 这里假设 lrcFile 是 []byte 类型，直接将其转换为字符串
		content := string(lrcFile) // 将字节数组转换为字符串
		song.Lyric = content
	}

	if err := s.songDao.UpdateSongLrc(song); err != nil {
		return "更新失败", err
	}

	return "更新成功", nil
}

package dto

type ListSongRequest struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"` // 主键，自增
	SongID     int    `json:"songId" gorm:"not null"`             // 歌曲ID，非空
	SongListID string `json:"songListId" gorm:"not null"`         // 歌单ID，非空
}

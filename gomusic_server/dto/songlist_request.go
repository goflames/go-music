package dto

type SongListRequest struct {
	ID           int    `json:"id" gorm:"primaryKey;autoIncrement"` // 主键，自增
	Title        string `json:"title" gorm:"size:255;not null"`     // 歌单标题，非空
	Pic          string `json:"pic" gorm:"size:255"`                // 歌单图片，可以为空
	Style        string `json:"style" gorm:"size:255"`              // 歌单风格，可以为空
	Introduction string `json:"introduction" gorm:"size:1024"`      // 歌单介绍，可以为空
}

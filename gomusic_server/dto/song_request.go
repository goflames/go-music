package dto

type SongRequest struct {
	ID           int    `json:"id" gorm:"column:id"`
	SingerID     int    `json:"singerId" gorm:"column:singer_id"`
	Name         string `json:"name" gorm:"column:name"`
	Introduction string `json:"introduction" gorm:"column:introduction"`
	CreateTime   string `json:"createTime" gorm:"column:create_time"`
	UpdateTime   string `json:"updateTime" gorm:"column:update_time"`
	Pic          string `json:"pic" gorm:"column:pic"`
	Lyric        string `json:"lyric" gorm:"column:lyric"`
	URL          string `json:"url" gorm:"column:url"`
}

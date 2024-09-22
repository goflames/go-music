package models

type SongList struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title        string `gorm:"type:varchar(255)" json:"title"`
	Pic          string `gorm:"type:varchar(255)" json:"pic"`
	Style        string `gorm:"type:varchar(255)" json:"style"`
	Introduction string `gorm:"type:text" json:"introduction"`
}

func (SongList) TableName() string {
	return "song_list"
}

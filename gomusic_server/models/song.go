package models

import "time"

type Song struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	SingerID     int       `gorm:"not null" json:"singerId"`
	Name         string    `gorm:"type:varchar(45);not null" json:"name"`
	Introduction string    `gorm:"type:varchar(255)" json:"introduction"`
	CreateTime   time.Time `gorm:"type:datetime;not null" json:"createTime"`
	UpdateTime   time.Time `gorm:"type:datetime;not null" json:"updateTime"`
	Pic          string    `gorm:"type:varchar(255)" json:"pic"`
	Lyric        string    `gorm:"type:text" json:"lyric"`
	URL          string    `gorm:"type:varchar(255);not null" json:"url"`
}

func (Song) TableName() string {
	return "song"
}

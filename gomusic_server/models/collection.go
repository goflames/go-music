package models

import "time"

type Collection struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     int       `json:"userId"`
	Type       int       `json:"type"`
	SongID     int       `json:"songId"`
	SongListID int       `json:"songListId"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
}

// TableName sets the table name for this struct
func (Collection) TableName() string {
	return "collect"
}

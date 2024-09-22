package models

import "time"

type Comment struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`        // 主键，自动递增
	UserID     uint      `gorm:"not null" json:"userId"`                    // 用户ID
	SongID     uint      `gorm:"not null" json:"songId"`                    // 歌曲ID
	SongListID uint      `gorm:"not null" json:"songListId"`                // 歌单ID
	Content    string    `gorm:"type:varchar(255);not null" json:"content"` // 评论内容
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`          // 创建时间，自动生成
	Type       byte      `gorm:"type:tinyint;not null" json:"type"`         // 评论类型
	Up         int       `gorm:"type:int" json:"up"`                        // 点赞数
}

func (Comment) TableName() string {
	return "comment"
}

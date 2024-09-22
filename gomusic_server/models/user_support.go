package models

// UserSupport 表示 user_support 表的结构体
type UserSupport struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CommentID int    `gorm:"not null" json:"commentId"`
	UserID    string `gorm:"type:varchar(45);not null" json:"userId"`
}

// TableName 设置模型对应的表名
func (UserSupport) TableName() string {
	return "user_support"
}

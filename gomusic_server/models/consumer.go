package models

import "time"

// Consumer represents the consumer model in the database
type Consumer struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string     `gorm:"type:varchar(255);not null" json:"username"`
	Password     string     `gorm:"type:varchar(255);not null" json:"password"`
	Sex          *byte      `gorm:"type:tinyint(1)" json:"sex"` // Use *byte for nullable Byte
	PhoneNum     string     `gorm:"type:varchar(20)" json:"phoneNum"`
	Email        string     `gorm:"type:varchar(255)" json:"email"`
	Birth        *time.Time `gorm:"type:date" json:"birth"` // Use *time.Time for nullable Date
	Introduction string     `gorm:"type:text" json:"introduction"`
	Location     string     `gorm:"type:varchar(255)" json:"location"`
	Avator       string     `gorm:"type:varchar(255)" json:"avator"`
	CreateTime   time.Time  `gorm:"autoCreateTime" json:"createTime"`
	UpdateTime   time.Time  `gorm:"autoUpdateTime" json:"updateTime"`
}

// TableName overrides the default table name
func (Consumer) TableName() string {
	return "consumer"
}

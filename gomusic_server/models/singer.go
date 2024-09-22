package models

import (
	"time"
)

type Singer struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"type:varchar(255)" json:"name"`
	Sex          *byte     `gorm:"type:tinyint" json:"sex"` // Use pointer to handle null values
	Pic          string    `gorm:"type:varchar(255)" json:"pic"`
	Birth        time.Time `gorm:"type:date" json:"birth"` // Use pointer to handle null values
	Location     string    `gorm:"type:varchar(255)" json:"location"`
	Introduction string    `gorm:"type:text" json:"introduction"`
}

func (Singer) TableName() string {
	return "singer"
}

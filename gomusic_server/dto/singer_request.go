package dto

type SingerRequest struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"` // 主键，自增
	Name         string `gorm:"size:255;not null" json:"name"`      // 字符串类型，最大长度255，不允许为null
	Sex          *byte  `gorm:"type:tinyint" json:"sex"`            // 字节类型，对应数据库中的 tinyint
	Pic          string `gorm:"size:255" json:"pic"`                // 字符串类型，最大长度255
	Birth        string `gorm:"type:date" json:"birth"`             // 时间类型，对应数据库中的日期
	Location     string `gorm:"size:255" json:"location"`           // 字符串类型，最大长度255
	Introduction string `gorm:"size:255" json:"introduction"`       // 字符串类型，最大长度255
}

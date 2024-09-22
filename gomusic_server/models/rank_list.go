package models

type RankList struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement;not null"` // 自增主键
	SongListID uint64 `gorm:"not null"`                          // song_list_id
	ConsumerID uint64 `gorm:"not null;uniqueIndex:consumerId"`   // consumer_id, 唯一索引
	Score      uint32 `gorm:"not null;default:0"`                // score, 默认为0
}

// TableName 设置表名为 rank_list
func (RankList) TableName() string {
	return "rank_list"
}

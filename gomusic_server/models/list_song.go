package models

// ListSong represents the `list_song` table structure with a unique index
type ListSong struct {
	ID         uint `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	SongID     uint `gorm:"not null;uniqueIndex:idx_song_list;column:song_id" json:"songId"`
	SongListID uint `gorm:"not null;uniqueIndex:idx_song_list;column:song_list_id" json:"songListId"`
}

// TableName sets the insert table name for this struct type
func (ListSong) TableName() string {
	return "list_song"
}

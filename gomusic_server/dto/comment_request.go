package dto

import (
	"gomusic_server/models" // 请根据你的项目结构调整导入路径
	"strconv"
	"time"
)

type CommentRequest struct {
	ID         int       `json:"id"`
	UserID     int       `json:"userId"`
	SongID     int       `json:"songId"`
	SongListID string    `json:"songListId"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"createTime"`
	NowType    byte      `json:"nowType"`
	Up         int       `json:"up"`
}

func (cr CommentRequest) ToComment() models.Comment {
	return models.Comment{
		ID:         uint(cr.ID),                      // 将 int 转换为 uint
		UserID:     uint(cr.UserID),                  // 将 string 转换为 uint
		SongID:     uint(cr.SongID),                  // 将 string 转换为 uint
		SongListID: parseStringToUint(cr.SongListID), // 将 string 转换为 uint
		Content:    cr.Content,                       // string 类型直接赋值
		CreateTime: cr.CreateTime,                    // time.Time 类型直接赋值
		Type:       cr.NowType,                       // byte 类型直接赋值
		Up:         cr.Up,                            // int 类型直接赋值
	}
}

// parseStringToUint 将字符串转换为 uint
func parseStringToUint(s string) uint {
	// 将字符串解析为 uint64
	result, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0 // 如果解析失败，返回 0
	}
	// 将 uint64 转换为 uint 类型
	return uint(result)
}

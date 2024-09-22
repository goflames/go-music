package dto

import (
	"gomusic_server/models"
	"strconv"
	"time"
)

type CollectRequest struct {
	ID         uint8      `json:"id,omitempty"`
	UserID     int8       `json:"userId,omitempty"`
	Type       string     `json:"type,omitempty"`
	SongID     int        `json:"songId,omitempty"`
	SongListID string     `json:"songListId,omitempty"`
	CreateTime *time.Time `json:"createTime,omitempty"`
}

func (req *CollectRequest) ToCollect() models.Collection {

	// 处理 CreateTime 的默认值
	var createTime time.Time
	if req.CreateTime != nil {
		createTime = *req.CreateTime
	} else {
		createTime = time.Now() // 如果没有提供，使用当前时间
	}
	//songId, _ := strconv.Atoi(req.SongID)
	songType, _ := strconv.Atoi(req.Type)
	songListID, _ := strconv.Atoi(req.SongListID)

	return models.Collection{
		ID:         int(req.ID),     // uint8 转换为 int
		UserID:     int(req.UserID), // int8 转换为 int
		Type:       songType,        // string 转换为 int
		SongID:     req.SongID,      // int8 转换为 int
		SongListID: songListID,      // int8 转换为 int
		CreateTime: createTime,      // 转换后的时间
	}
}

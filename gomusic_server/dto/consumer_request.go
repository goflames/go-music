package dto

import (
	"gomusic_server/models"
	"strings"
	"time"
)

type ConsumerRequest struct {
	Id           uint       `json:"id"`
	Username     string     `json:"username"`
	OldPassword  string     `json:"oldPassword"` // 旧密码
	Password     string     `json:"password"`
	Sex          *byte      `json:"sex"`
	PhoneNum     string     `json:"phoneNum"`
	Email        string     `json:"email"`
	Birth        *time.Time `json:"birth"`
	Introduction string     `json:"introduction"`
	Location     string     `json:"location"`
	Avator       string     `json:"avator"`
	CreateTime   time.Time  `json:"createTime"`
}

func (r *ConsumerRequest) ToConsumer() *models.Consumer {
	return &models.Consumer{
		ID:           r.Id,
		Username:     r.Username,
		Password:     strings.TrimSpace(r.Password), // 去除密码中的空格
		Sex:          r.Sex,
		PhoneNum:     r.PhoneNum,
		Email:        r.Email,
		Birth:        r.Birth,
		Introduction: r.Introduction,
		Location:     r.Location,
		Avator:       r.Avator,
		CreateTime:   time.Now(),
	}
}

package dto

import (
	"gomusic_server/models"
	"strconv"
)

type UserSupportRequest struct {
	ID        int `json:"id"`
	CommentID int `json:"commentId"`
	UserID    int `json:"userId"`
}

// ToModel 将 UserSupportRequest 转换为 UserSupport
func (r *UserSupportRequest) ToModel() *models.UserSupport {
	return &models.UserSupport{
		ID:        r.ID,
		CommentID: r.CommentID,
		UserID:    strconv.Itoa(r.UserID), // Convert int to string
	}
}

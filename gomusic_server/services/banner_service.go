package service

import (
	"github.com/gin-gonic/gin"
	"gomusic_server/dao"
	"gomusic_server/models"
	"gorm.io/gorm"
)

type BannerService struct {
	BannerDAO *dao.BannerDAO
}

func NewBannerService(db *gorm.DB) *BannerService {
	return &BannerService{dao.NewBannerDAO(db)}
}

func (s *BannerService) GetAllBanners(ctx *gin.Context) ([]models.Banner, error) {
	banners, err := s.BannerDAO.GetAllBanners()
	if err != nil {
		return nil, err
	}
	return banners, nil
}

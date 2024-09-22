package dao

import (
	"gomusic_server/models"
	"gorm.io/gorm"
)

type BannerDAO struct {
	Db *gorm.DB
}

func NewBannerDAO(db *gorm.DB) *BannerDAO {
	return &BannerDAO{Db: db}
}

func (dao *BannerDAO) GetAllBanners() ([]models.Banner, error) {
	var banners []models.Banner
	if err := dao.Db.Find(&banners).Error; err != nil {
		return nil, err
	}
	return banners, nil
}

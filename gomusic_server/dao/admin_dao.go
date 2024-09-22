package dao

import (
	"errors"
	"gomusic_server/models"
	"gorm.io/gorm"
)

type AdminDAO struct {
	db *gorm.DB
}

func NewAdminDAO(db *gorm.DB) *AdminDAO {
	return &AdminDAO{db: db}
}

func (dao *AdminDAO) GetAdmin(name string) (*models.Admin, error) {
	var admin *models.Admin
	result := dao.db.Where("name = ?", name).First(&admin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 记录未找到的情况
			return nil, nil
		}
		// 其他错误
		return nil, result.Error
	}
	return admin, result.Error
}

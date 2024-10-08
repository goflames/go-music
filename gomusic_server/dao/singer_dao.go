package dao

import (
	"gomusic_server/common"
	"gomusic_server/dto"
	"gomusic_server/models"
	"gorm.io/gorm"
	"time"
)

type SingerDAO struct {
	Db *gorm.DB
}

func NewSingerDAO(db *gorm.DB) *SingerDAO {
	return &SingerDAO{Db: db}
}

func (dao *SingerDAO) GetAllSingers() ([]models.Singer, error) {
	var singers []models.Singer
	err := dao.Db.Find(&singers).Error
	if err != nil {
		return nil, err
	}
	return singers, err
}

func (dao *SingerDAO) GetSingersByGender(gender int) ([]models.Singer, error) {
	var singers []models.Singer
	err := dao.Db.Where("sex = ?", gender).Find(&singers).Error
	return singers, err
}

func (dao *SingerDAO) UpdateSingerInfo(request dto.SingerRequest) error {
	var singer models.Singer
	db := dao.Db.First(&singer, request.ID)
	singer.Name = request.Name
	singer.Sex = request.Sex
	parsedTime, err := time.Parse("2006-01-02", request.Birth)
	singer.Birth = parsedTime // 需要将字符串转换为 time.Time
	if err != nil {
		return err
	}
	singer.Location = request.Location
	singer.Introduction = request.Introduction

	return db.Save(&singer).Omit("pic").Error
}

func (dao *SingerDAO) UpdateSingerImg(singerId int, pic string) common.Response {
	tx := dao.Db.Model(&models.Singer{}).Where("id = ?", singerId).Update("pic", pic)
	if tx.Error != nil {
		return common.Error("数据库更新歌手图片失败")
	}
	return common.SuccessWithData("更新成功", pic)
}

func (dao *SingerDAO) DeleteSingerById(singerId int) common.Response {
	tx := dao.Db.Delete(&models.Singer{}, singerId)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return common.Error("删除歌手失败...")
	}
	return common.Success("删除歌手成功！")
}

func (dao *SingerDAO) AddSinger(singer models.Singer) common.Response {
	tx := dao.Db.Create(&singer)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return common.Error("新增歌手信息失败！")
	}
	return common.Success("新增歌手成功！")
}

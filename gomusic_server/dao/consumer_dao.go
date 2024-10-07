package dao

import (
	"gomusic_server/common"
	"gomusic_server/models"
	"gorm.io/gorm"
	"log"
)

type ConsumerDAO struct {
	db *gorm.DB
}

func NewConsumerDAO(db *gorm.DB) *ConsumerDAO {
	return &ConsumerDAO{
		db: db,
	}
}

func (dao *ConsumerDAO) GetUserById(id int8) (models.Consumer, error) {
	var consumer models.Consumer
	err := dao.db.First(&consumer, id).Error
	return consumer, err
}

func (dao *ConsumerDAO) GetUserByUsername(username string) (*models.Consumer, error) {
	var consumer *models.Consumer
	err := dao.db.Where("username = ?", username).Find(&consumer).Error
	return consumer, err
}

func (dao *ConsumerDAO) GetUserByEmail(email string) (*models.Consumer, error) {
	var consumer *models.Consumer
	err := dao.db.Where("email = ?", email).Find(&consumer).Error
	return consumer, err
}

func (dao *ConsumerDAO) CheckEmailDuplicate(email string) (bool, error) {
	var consumer *models.Consumer
	tx := dao.db.Where("email = ?", email).Find(&consumer)
	return tx.RowsAffected == 1, tx.Error
}

func (dao *ConsumerDAO) Create(consumer *models.Consumer) (*models.Consumer, error) {
	tx := dao.db.Create(consumer)
	if tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return consumer, tx.Error
}

func (dao *ConsumerDAO) UpdatePassword(consumer *models.Consumer) (int64, error) {
	tx := dao.db.Save(&consumer)
	return tx.RowsAffected, tx.Error
}

func (dao *ConsumerDAO) Update(consumer *models.Consumer) (int64, error) {
	tx := dao.db.Model(&consumer).Omit("Password").Updates(consumer)
	return tx.RowsAffected, tx.Error
}

func (dao *ConsumerDAO) DeleteById(userId int8) bool {
	tx := dao.db.Delete(&models.Consumer{}, userId)
	return tx.RowsAffected > 0
}

func (dao *ConsumerDAO) GetAllUser() []models.Consumer {
	var users []models.Consumer
	tx := dao.db.Find(&users)
	if tx.Error != nil {
		return nil
	}
	return users
}

func (dao *ConsumerDAO) UpdateUserImg(userId int8, pic string) common.Response {
	log.Printf("进入用户头像更新数据库方法")
	tx := dao.db.Model(&models.Consumer{}).Where("id = ?", userId).Update("avator", pic)
	if tx.Error != nil || tx.RowsAffected < 1 {
		log.Print("用户头像更新失败.....")
		return common.Error("数据库更新失败")
	}
	return common.SuccessWithData("更新成功", pic)
}

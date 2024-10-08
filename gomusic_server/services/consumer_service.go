package service

import (
	"gomusic_server/common"
	"gomusic_server/dao"
	"gomusic_server/dto"
	"gomusic_server/models"
	"gomusic_server/utils"
	"gorm.io/gorm"
	"strings"
)

type ConsumerService struct {
	consumerDAO *dao.ConsumerDAO
}

func NewConsumerService(db *gorm.DB) *ConsumerService {
	return &ConsumerService{dao.NewConsumerDAO(db)}
}

func (s *ConsumerService) GetUserById(id int) (models.Consumer, error) {
	consumer, err := s.consumerDAO.GetUserById(id)
	return consumer, err
}

func (s *ConsumerService) LoginStatus(loginRequest dto.ConsumerRequest) (*models.Consumer, bool, error) {
	username := loginRequest.Username
	password := loginRequest.Password

	consumer, err := s.consumerDAO.GetUserByUsername(username)
	if err != nil || consumer == nil {
		return nil, false, err
	}

	checkPassword := utils.CheckPassword(strings.TrimSpace(consumer.Password), strings.TrimSpace(password))
	return consumer, checkPassword, err
}

func (s *ConsumerService) EmailStatus(loginRequest dto.ConsumerRequest) (*models.Consumer, bool, error) {
	email := loginRequest.Email
	password := loginRequest.Password

	consumer, err := s.consumerDAO.GetUserByEmail(email)
	if err != nil || consumer == nil {
		return nil, false, err
	}

	checkPassword := utils.CheckPassword(strings.TrimSpace(consumer.Password), strings.TrimSpace(password))
	return consumer, checkPassword, err
}

func (s *ConsumerService) AddUser(registryRequest *dto.ConsumerRequest) (common.Response, error) {
	// 检查用户名是否重复
	existingUser, err := s.consumerDAO.GetUserByUsername(registryRequest.Username)
	if err != nil {
		return common.Fatal("数据库查询错误"), err
	}
	if existingUser.Username != "" {
		return common.Warning("用户名已注册"), nil
	}

	consumer := &models.Consumer{}
	*consumer = *registryRequest.ToConsumer()

	// 密码加密
	consumer.Password, err = utils.HashPassword(consumer.Password)
	if err != nil {
		return common.Fatal("密码加密失败！"), err
	}

	// 处理空值
	if strings.TrimSpace(consumer.PhoneNum) == "" {
		return common.Error("手机号不能为空"), err

	}
	if consumer.Email == "" {
		return common.Error("邮箱不能为空"), err
	}
	consumer.Avator = "/img/avatorImages/user.jpg"

	// 检查邮箱是否重复
	if registryRequest.Email != "" {
		duplicate, _ := s.consumerDAO.CheckEmailDuplicate(registryRequest.Email)
		if duplicate {
			return common.Fatal("该邮箱已注册"), nil
		}
	}

	// 插入用户
	create, err := s.consumerDAO.Create(consumer)
	if create == nil {
		return common.Error("注册失败"), err

	}
	return common.Success("注册成功"), nil
}

func (s *ConsumerService) UpdateUser(updateRequest dto.ConsumerRequest) (common.Response, error) {
	oldinfo, _ := s.consumerDAO.GetUserById(int(updateRequest.Id))
	if oldinfo.Username != updateRequest.Username {
		existingUser, err := s.consumerDAO.GetUserByUsername(updateRequest.Username)
		if err != nil {
			return common.Fatal("数据库查询错误"), err
		}
		if existingUser.Username != "" {
			return common.Warning("用户名已注册"), nil
		}
	}

	if oldinfo.Email != updateRequest.Email {
		// 检查邮箱是否重复
		duplicate, _ := s.consumerDAO.CheckEmailDuplicate(updateRequest.Email)
		if duplicate {
			return common.Fatal("该邮箱已注册"), nil
		}
	}

	consumer := &models.Consumer{}
	*consumer = *updateRequest.ToConsumer()

	row, err := s.consumerDAO.Update(consumer)
	if row < 1 {
		return common.Error("更新失败"), err
	}
	return common.Success("更新成功"), nil
}

func (s *ConsumerService) UpdatePassword(consumer models.Consumer) (int64, error) {
	count, err := s.consumerDAO.UpdatePassword(&consumer)
	return count, err
}

func (s *ConsumerService) DeleteById(userId int) bool {
	return s.consumerDAO.DeleteById(userId)
}

func (s *ConsumerService) GetAllUser() []models.Consumer {
	return s.consumerDAO.GetAllUser()
}

func (s *ConsumerService) UpdateUserImg(userId int, pic string) common.Response {
	return s.consumerDAO.UpdateUserImg(userId, pic)
}

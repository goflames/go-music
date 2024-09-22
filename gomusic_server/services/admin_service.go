package service

import (
	"gomusic_server/dao"
	"gorm.io/gorm"
)

type AdminService struct {
	adminDAO *dao.AdminDAO
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{dao.NewAdminDAO(db)}
}

func (s *AdminService) AdminLoginStatus(name string, password string) bool {
	admin, err := s.adminDAO.GetAdmin(name)
	if err != nil {
		return false
	}
	return admin.Password == password
}

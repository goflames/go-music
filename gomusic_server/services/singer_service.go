package service

import (
	"gomusic_server/common"
	"gomusic_server/dao"
	"gomusic_server/dto"
	"gomusic_server/models"
	"gorm.io/gorm"
)

type SingerService struct {
	SingerDAO *dao.SingerDAO
}

func NewSingerService(db *gorm.DB) *SingerService {
	return &SingerService{dao.NewSingerDAO(db)}
}

func (s *SingerService) GetAllSingers() ([]models.Singer, error) {
	singers, err := s.SingerDAO.GetAllSingers()
	if err != nil {
		return nil, err
	}
	return singers, err
}

func (s *SingerService) GetSingersByGender(gender int8) ([]models.Singer, error) {
	singers, err := s.SingerDAO.GetSingersByGender(gender)
	if err != nil {
		return nil, err
	}
	return singers, err
}

func (s *SingerService) UpdateSingerInfo(request dto.SingerRequest) error {
	return s.SingerDAO.UpdateSingerInfo(request)
}

func (s *SingerService) UpdateSingerImg(singerId int8, pic string) common.Response {
	return s.SingerDAO.UpdateSingerImg(singerId, pic)
}

func (s *SingerService) DeleteSingerById(singerId int8) common.Response {
	return s.SingerDAO.DeleteSingerById(singerId)
}

func (s *SingerService) AddSinger(singer models.Singer) common.Response {
	return s.SingerDAO.AddSinger(singer)
}

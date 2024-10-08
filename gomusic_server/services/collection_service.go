package service

import (
	"gomusic_server/common"
	"gomusic_server/dao"
	"gomusic_server/models"
	"gorm.io/gorm"
)

type CollectionService struct {
	collectionDAO *dao.CollectionDAO
}

func NewCollectionService(db *gorm.DB) *CollectionService {
	return &CollectionService{dao.NewCollectionDAO(db)}
}

func (s *CollectionService) GetCollectionByUserId(userId int) ([]*models.Collection, error) {
	collections, err := s.collectionDAO.GetCollectionByUserId(userId)
	return collections, err
}

func (s *CollectionService) ExistSongID(collect models.Collection) (bool, error) {
	count, err := s.collectionDAO.ExistSongID(collect.UserID, collect.SongID)
	return count > 0, err
}

func (s *CollectionService) AddCollection(collect models.Collection) common.Response {
	return s.collectionDAO.AddCollection(collect)
}

func (s *CollectionService) DeleteCollection(userId, songId int) common.Response {
	return s.collectionDAO.DeleteCollection(userId, songId)
}

package service

import (
	"market/internal/repository"
	"market/models"
)

type InfoService struct {
	repo repository.Information
}

func NewInfoService(repo repository.Information) *InfoService {
	return &InfoService{repo: repo}
}

func (s *InfoService) GetInfo(ID int) (models.Info, error) {
	return s.repo.GetInfo(ID)

}

package main

import (
	"fmt"
	"vk/repository"
)

type CountUsersByAuthor struct {
	AuthorID   int
	CountUsers int
}

type AnalyticService struct {
	repo         repository.Repository
	analyticDate *AnalyticDate
}

func NewAnalyticService(repository repository.Repository, analyticDate *AnalyticDate) *AnalyticService {
	return &AnalyticService{repo: repository, analyticDate: analyticDate}
}

func (s *AnalyticService) Save(authorID int, userID int) error {
	dto := repository.AnalyticDTO{
		UserID:    userID,
		AuthorID:  authorID,
		CreatedAt: s.analyticDate.CurrentDate(),
	}

	err := s.repo.Save(dto)
	if err != nil {
		return fmt.Errorf("save error: %v", err)
	}

	return nil
}

func (s *AnalyticService) List(authorIDs []int) ([]CountUsersByAuthor, error) {
	result := make([]CountUsersByAuthor, 0)

	items, err := s.repo.CountUsersByAuthor(authorIDs, s.analyticDate.PrevDate())
	if err != nil {
		return result, fmt.Errorf("analyticService->List error: %v", err)
	}

	for authorID, countUsers := range items {
		result = append(result, CountUsersByAuthor{AuthorID: authorID, CountUsers: countUsers})
	}

	return result, nil
}

func (s *AnalyticService) DeleteOld() error {
	date := s.analyticDate.PrevDate().AddDate(0, 0, -1)
	_, err := s.repo.DeleteByDate(date)

	if err != nil {
		return fmt.Errorf("analyticService->DeleteOld error: %v", err)
	}

	return nil
}

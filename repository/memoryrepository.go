package repository

import (
	"time"
)

type MemoryRepository struct {
	items map[string]map[int]map[int]int
}

func NewMemoryRepository() Repository {
	return &MemoryRepository{
		items: make(map[string]map[int]map[int]int),
	}
}

func (m *MemoryRepository) Save(dto AnalyticDTO) error {
	if _, ok := m.items[m.formatDate(dto.CreatedAt)]; !ok {
		m.items[m.formatDate(dto.CreatedAt)] = make(map[int]map[int]int)
	}

	if _, ok := m.items[m.formatDate(dto.CreatedAt)][dto.AuthorID]; !ok {
		m.items[m.formatDate(dto.CreatedAt)][dto.AuthorID] = make(map[int]int)
	}

	m.items[m.formatDate(dto.CreatedAt)][dto.AuthorID][dto.UserID] = dto.UserID

	return nil
}

func (m *MemoryRepository) CountUsersByAuthor(authorIDs []int, date time.Time) (map[int]int, error) {
	result := make(map[int]int)

	authors, ok := m.items[m.formatDate(date)]
	if !ok {
		return result, nil
	}

	for _, authorID := range authorIDs {
		users, ok := authors[authorID]
		if !ok {
			result[authorID] = 0

			continue
		}

		result[authorID] = len(users)
	}

	return result, nil
}

func (m *MemoryRepository) DeleteByDate(date time.Time) (bool, error) {
	formatDate := m.formatDate(date)

	if _, ok := m.items[formatDate]; !ok {
		return false, nil
	}

	delete(m.items, formatDate)

	return true, nil
}

func (m *MemoryRepository) formatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

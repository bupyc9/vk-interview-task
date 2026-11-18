package repository

import (
	"time"
)

type AnalyticDTO struct {
	UserID    int
	AuthorID  int
	CreatedAt time.Time
}

type Repository interface {
	Save(dto AnalyticDTO) error
	CountUsersByAuthor(authorIDs []int, date time.Time) (map[int]int, error)
	DeleteByDate(date time.Time) (bool, error)
}

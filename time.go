package main

import "time"

type AnalyticDate struct {
	now time.Time
}

func NewAnalyticDate(now time.Time) *AnalyticDate {
	return &AnalyticDate{now: now}
}

func (t *AnalyticDate) CurrentDate() time.Time {
	return t.now
}

func (t *AnalyticDate) PrevDate() time.Time {
	return t.now.AddDate(0, 0, -1)
}

func (t *AnalyticDate) SetNow(now time.Time) {
	t.now = now
}

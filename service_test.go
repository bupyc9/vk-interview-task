package main

import (
	"testing"
	"time"
	"vk/repository"
)

func TestAnalyticService_Save(t *testing.T) {
	date := NewAnalyticDate(time.Now())

	repo := repository.NewMemoryRepository()
	service := NewAnalyticService(repo, date)

	service.Save(1, 1)
	service.Save(1, 1)
	service.Save(1, 2)
	service.Save(2, 1)

	date.SetNow(time.Now().AddDate(0, 0, 1))

	got, err := service.List([]int{1, 2, 3})

	if err != nil {
		t.Error(err)
	}

	if len(got) != 3 {
		t.Errorf("count = %v; expected = %v", len(got), 3)
	}

	if got[0].CountUsers != 2 {
		t.Errorf("uID = %v; got = %v; expected = %v", 1, got[0].CountUsers, 2)
	}
	if got[1].CountUsers != 1 {
		t.Errorf("uID = %v; got = %v; expected = %v", 1, got[1].CountUsers, 1)
	}
	if got[2].CountUsers != 0 {
		t.Errorf("uID = %v; got = %v; expected = %v", 1, got[2].CountUsers, 0)
	}
}

func TestAnalyticService_DeleteOld(t *testing.T) {
	date := NewAnalyticDate(time.Now().AddDate(0, 0, -2))

	repo := repository.NewMemoryRepository()
	service := NewAnalyticService(repo, date)

	service.Save(1, 1)
	service.Save(2, 1)

	date.SetNow(time.Now().AddDate(0, 0, -1))

	service.Save(1, 1)
	service.Save(1, 2)
	service.Save(2, 1)
	service.Save(2, 2)

	date.SetNow(time.Now())

	service.Save(1, 1)
	service.Save(1, 2)
	service.Save(1, 3)
	service.Save(2, 1)
	service.Save(2, 2)
	service.Save(2, 3)

	err := service.DeleteOld()
	if err != nil {
		t.Errorf("deleteOld error: %v", err)
	}

	date.SetNow(time.Now().AddDate(0, 0, -1))
	got, err := service.List([]int{1, 2})

	if len(got) != 0 {
		t.Errorf("count = %v; expected = %v", len(got), 0)
	}

	date.SetNow(time.Now())
	got, err = service.List([]int{1, 2})

	if len(got) != 2 {
		t.Errorf("count = %v; expected = %v", len(got), 2)
	}
	if got[0].CountUsers != 2 {
		t.Errorf("uID = %v; got = %v; expected = %v", 1, got[0].CountUsers, 2)
	}
	if got[1].CountUsers != 2 {
		t.Errorf("uID = %v; got = %v; expected = %v", 1, got[1].CountUsers, 2)
	}

	date.SetNow(time.Now().AddDate(0, 0, +1))
	got, err = service.List([]int{1, 2})

	if len(got) != 2 {
		t.Errorf("count = %v; expected = %v", len(got), 2)
	}
	if got[0].CountUsers != 3 {
		t.Errorf("uID = %v; got = %v; expected = %v", 1, got[0].CountUsers, 3)
	}
	if got[1].CountUsers != 3 {
		t.Errorf("uID = %v; got = %v; expected = %v", 1, got[1].CountUsers, 3)
	}
}

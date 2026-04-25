package service

import "github.com/revpos/eTrak/internal/repository"

type EventService struct {
	Repo *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
	return &EventService{Repo: repo}
}

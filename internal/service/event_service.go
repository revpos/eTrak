package service

import (
	"context"

	"github.com/revpos/eTrak/internal/models"
	"github.com/revpos/eTrak/internal/repository"
)

type EventService struct {
	Repo *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
	return &EventService{Repo: repo}
}

func (s *EventService) TrackEvent(ctx context.Context, e *models.Event) error {
	// later: validation, enrichment, async queue
	return s.Repo.InsertEvent(ctx, e)
}

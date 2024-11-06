package service

import (
	"context"
	"fmt"
	"time"

	"github.com/zuekveer/BWA/internal/entity"
)

type repository interface {
	Create(ctx context.Context, user entity.Event) error
	Update(ctx context.Context, userID entity.EventId, eventChanges entity.EventChanges) error
	Get(ctx context.Context, from time.Time, to time.Time) ([]entity.Event, error)
}

type EventService struct {
	repo repository
}

func NewEventService(repo repository) *EventService {
	return &EventService{repo}
}

func (s *EventService) Create(ctx context.Context, event entity.Event) (entity.Event, error) {
	var zero entity.Event

	ev := entity.NewEvent(event)
	err := s.repo.Create(ctx, ev)
	if err != nil {
		return zero, fmt.Errorf("failed to persist new event: %w", err)
	}
	return ev, nil
}

func (s *EventService) Update(ctx context.Context, eventId entity.EventId, changes entity.EventChanges) error {
	return s.repo.Update(ctx, eventId, changes)
}

func (s *EventService) Get(ctx context.Context, from time.Time, to time.Time) ([]entity.Event, error) {
	return s.repo.Get(ctx, from, to)
}

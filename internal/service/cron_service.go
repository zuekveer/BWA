package service

import (
	"context"
	"fmt"
	"time"

	"github.com/zuekveer/BWA/internal/entity"
	"github.com/zuekveer/BWA/internal/logger"
)

type cronRepository interface {
	CleanEvents(ctx context.Context, time time.Time) (err error)
	GetNotNotified(ctx context.Context) (events []entity.Event, err error)
	SetNotificationDate(ctx context.Context, eventId entity.EventId, time time.Time) (err error)
}

type CronService struct {
	repo cronRepository
	l    logger.Logger
}

func NewCronService(repo cronRepository, l logger.Logger) *CronService {
	return &CronService{repo, l}
}

func (c *CronService) Notify(ctx context.Context) error {
	entries, err := c.repo.GetNotNotified(ctx)
	if err != nil {
		return fmt.Errorf("failed to get events for notify: %w", err)
	}
	now := time.Now()
	for _, v := range entries {
		start := v.StartDate
		end := v.EndDate
		if now.Before(end) && end.After(start) {
			before := now.Add(-time.Duration(v.DelayedTime) * time.Second)
			if err := c.repo.SetNotificationDate(ctx, v.ID, before); err != nil {
				return fmt.Errorf("failed to sendNotification: %w", err)
			}
			c.l.Infof("Send notification with id:%s", v.ID)
		}
	}
	return nil
}

func (c *CronService) Clean(ctx context.Context) error {
	t := time.Now().AddDate(-1, 0, 0)
	if err := c.repo.CleanEvents(ctx, t); err != nil {
		return fmt.Errorf("failed to clean events: %w", err)
	}
	return nil
}

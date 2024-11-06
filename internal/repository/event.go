package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/zuekveer/BWA/internal/databases"
	"github.com/zuekveer/BWA/internal/entity"
)

const createQuery = `
INSERT INTO events (
            id, header, creation_date, start_date, end_date, description, user_id, delayed_time, notification_send_date
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

type Repository struct {
	db *databases.DB
}

func NewRepository(db *databases.DB) *Repository {
	return &Repository{db}
}

type changesBuilder entity.EventChanges
type EventList []entity.Event

func (r *Repository) Create(ctx context.Context, event entity.Event) error {
	_, err := r.db.ExecContext(ctx, createQuery,
		event.ID,
		event.Header,
		event.CreationDate,
		event.EndDate,
		event.Description,
		event.UserId,
		event.DelayedTime,
		event.NotificationSendDate,
	)
	if err != nil {
		return fmt.Errorf("failed to create event: %w", err)
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, eventId entity.EventId, changes entity.EventChanges) error {
	qb := squirrel.Update("events").Where(squirrel.Eq{"id": eventId})
	//// NOTE::
	// объявим тип внутри пакета repositories, чтобы добавить дополнительное поведение.
	// в пакете users мы не можем этого сделать, т.к. есть знание о маппинге полей.
	qb = qb.SetMap(changesBuilder(changes).ToMap())

	query, args, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, eventId entity.EventId) error {
	db := squirrel.Delete("events").Where(squirrel.Eq{"id": eventId})
	_, err := db.ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}
	return nil
}

func (r *Repository) Get(ctx context.Context, from time.Time, to time.Time) (events []entity.Event, err error) {
	qb := squirrel.Select("id", "header", "creation_date", "start_date", "end_date", "description", "user_id", "delayed_time", "notification_send_date").
		From("events").
		Where(squirrel.GtOrEq{"creation_date": from}).
		Where(squirrel.LtOrEq{"end_date": to})
	query, args, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build conditions: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to exec query: %w", err)
	}
	return r.get(rows), nil
}

func (r *Repository) get(rows *sql.Rows) (events []entity.Event) {
	entries := make([]entity.Event, 0)
	for rows.Next() {
		var entry entity.Event
		_ = rows.Scan(
			&entry.ID,
			&entry.Header,
			&entry.CreationDate,
			&entry.StartDate,
			&entry.EndDate,
			&entry.Description,
			&entry.UserId,
			&entry.DelayedTime,
			&entry.NotificationSendDate,
		)
		entries = append(entries, entry)
	}
	return entries
}

func (r *Repository) CleanEvents(ctx context.Context, time time.Time) (err error) {
	qb := squirrel.Delete("events").
		Where(squirrel.Lt{"creation_date": time})
	query, args, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("failed to clean events: %w", err)
	}
	_, err = r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}
	return nil
}

func (r *Repository) GetNotNotified(ctx context.Context) (events []entity.Event, err error) {
	qb := squirrel.Select("id", "header", "creation_date", "start_date", "end_date", "description", "user_id", "delayed_time", "notification_send_date").
		From("events").
		Where(squirrel.Eq{"notification_send_date": nil})
	query, args, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build conditions: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to exec query: %w", err)
	}
	return r.get(rows), nil
}

func (r *Repository) SetNotificationDate(ctx context.Context, eventId entity.EventId, time time.Time) (err error) {
	qb := squirrel.Update("events").Where(squirrel.Eq{"id": eventId})
	qb = qb.Set("notification_send_date", time)
	query, args, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}
	return nil
}

func (b changesBuilder) ToMap() map[string]any {
	fields := make(map[string]any)
	if b.Description != "" {
		fields["description"] = b.Description
	}
	if b.Header != "" {
		fields["header"] = b.Header
	}
	if !time.Time.IsZero(b.EndDate) {
		fields["end_date"] = b.EndDate
	}
	if !time.Time.IsZero(b.EndDate) {
		fields["start_date"] = b.StartDate
	}
	if b.UserId != "" {
		fields["user_id"] = b.UserId
	}
	return fields
}

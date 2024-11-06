package entity

import (
	"time"

	"github.com/google/uuid"
)

type EventId = string

type Event struct {
	ID                   EventId   `json:"id"`
	Header               string    `json:"header"`
	CreationDate         time.Time `json:"creationDate"`
	StartDate            time.Time `json:"startDate"`
	EndDate              time.Time `json:"endDate"`
	Description          string    `json:"description"`
	UserId               string    `json:"userId"`
	DelayedTime          int       `json:"delayedTime"`
	NotificationSendDate time.Time `json:"notificationSendDate"`
}

func NewEvent(event Event) Event {
	return Event{
		ID:                   uuid.New().String(),
		Header:               event.Header,
		CreationDate:         event.CreationDate,
		StartDate:            event.StartDate,
		EndDate:              event.EndDate,
		Description:          event.Description,
		UserId:               event.UserId,
		DelayedTime:          event.DelayedTime,
		NotificationSendDate: event.NotificationSendDate,
	}
}

type EventChanges struct {
	Description string    `json:"description"`
	Header      string    `json:"header"`
	EndDate     time.Time `json:"endDate"`
	StartDate   time.Time `json:"startDate"`
	UserId      string    `json:"userId"`
}

package events

import (
	"github.com/google/uuid"
	"time"
)

type IEvent interface {
	EventId() uuid.UUID
	EventType() string
	CreatedAt() time.Time
}

type Event struct {
	eventId   uuid.UUID
	eventType string
	createdAt time.Time
}

func NewEvent(eventId uuid.UUID, eventType string) *Event {
	return &Event{
		eventId:   eventId,
		eventType: eventType,
		createdAt: time.Now(),
	}
}

func (e *Event) EventId() uuid.UUID {
	return e.eventId
}

func (e *Event) EventType() string {
	return e.eventType
}

func (e *Event) CreatedAt() time.Time {
	return e.createdAt
}

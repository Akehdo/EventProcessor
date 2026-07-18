package repository

import (
	"EventProcessor/internal/domain"
	"sync"

	"github.com/google/uuid"
)

type EventRepository struct {
	events map[uuid.UUID]domain.Event
	mu     sync.RWMutex
}

func NewEventRepository() *EventRepository {
	return &EventRepository{
		events: make(map[uuid.UUID]domain.Event),
	}
}

func (r *EventRepository) Save(event domain.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.events[event.ID]; exists {
		return domain.ErrEventAlreadyExists
	}

	r.events[event.ID] = event

	return nil
}

func (r *EventRepository) List() ([]domain.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	events := make([]domain.Event, 0, len(r.events))

	for _, event := range r.events {
		events = append(events, event)
	}

	return events, nil
}

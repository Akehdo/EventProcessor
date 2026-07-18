package service

import (
	"EventProcessor/internal/domain"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type EventRepository interface {
	Save(event domain.Event) error
	List() ([]domain.Event, error)
}

type EventService struct {
	repo EventRepository
}

func NewEventService(repo EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) Create(
	payload json.RawMessage,
) (domain.Event, error) {
	if len(payload) == 0 {
		return domain.Event{}, domain.ErrPayloadRequired
	}

	if !json.Valid(payload) {
		return domain.Event{}, domain.ErrPayloadInvalid
	}

	if string(payload) == "null" {
		return domain.Event{}, domain.ErrPayloadNull
	}

	now := time.Now()

	event := domain.Event{
		ID:        uuid.New(),
		Status:    domain.StatusPending,
		Payload:   payload,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Save(event); err != nil {
		return domain.Event{}, err
	}

	return event, nil
}

func (s *EventService) List() ([]domain.Event, error) {
	return s.repo.List()
}

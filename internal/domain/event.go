package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID
	Status    Status
	Payload   json.RawMessage
	CreatedAt time.Time
	UpdatedAt time.Time
}

package dto

import "encoding/json"

type CreateEventRequest struct {
	Payload json.RawMessage `json:"payload" binding:"required"`
}

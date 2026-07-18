package domain

import "errors"

var (
	ErrEventAlreadyExists = errors.New("event already exists")
	ErrPayloadRequired    = errors.New("payload is required")
	ErrPayloadInvalid     = errors.New("payload must be valid JSON")
	ErrPayloadNull        = errors.New("payload cannot be null")
)

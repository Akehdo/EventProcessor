package handler

import (
	"EventProcessor/internal/domain"
	"EventProcessor/internal/handler/dto"
	"EventProcessor/internal/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventProcessor interface {
	Enqueue(id uuid.UUID)
}

type EventsHandler struct {
	service   *service.EventService
	processor EventProcessor
}

func NewEventsHandler(
	service *service.EventService,
	processor EventProcessor,
) *EventsHandler {
	return &EventsHandler{
		service:   service,
		processor: processor,
	}
}

func (h *EventsHandler) Create(ctx *gin.Context) {
	var request dto.CreateEventRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	event, err := h.service.Create(request.Payload)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPayloadRequired),
			errors.Is(err, domain.ErrPayloadInvalid),
			errors.Is(err, domain.ErrPayloadNull):

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

		case errors.Is(err, domain.ErrEventAlreadyExists):
			ctx.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
		}

		return
	}

	h.processor.Enqueue(event.ID)

	ctx.JSON(http.StatusCreated, event)
}

func (h *EventsHandler) List(ctx *gin.Context) {
	events, err := h.service.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

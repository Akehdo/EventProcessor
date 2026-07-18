package main

import (
	"EventProcessor/internal/handler"
	"EventProcessor/internal/repository"
	"EventProcessor/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	eventRepository := repository.NewEventRepository()
	eventService := service.NewEventService(eventRepository)
	eventsHandler := handler.NewEventsHandler(eventService)

	router := gin.Default()

	router.POST("/events", eventsHandler.Create)
	router.GET("/events", eventsHandler.List)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

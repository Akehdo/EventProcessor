package main

import (
	"EventProcessor/internal/handler"
	"EventProcessor/internal/processor"
	"EventProcessor/internal/repository"
	"EventProcessor/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	eventRepository := repository.NewEventRepository()
	eventService := service.NewEventService(eventRepository)

	eventProcessor := processor.NewEventProcessor(eventService, 100)
	eventProcessor.Start(3)
	defer eventProcessor.Stop()
	
	eventsHandler := handler.NewEventsHandler(eventService, eventProcessor)

	router := gin.Default()

	router.POST("/events", eventsHandler.Create)
	router.GET("/events", eventsHandler.List)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

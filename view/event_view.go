package view

import (
	"my-go-app/model"

	"github.com/gin-gonic/gin"
)

// ProcessEvents はイベントデータを加工してJSON化する
func ProcessEvents(events []model.Event) gin.H {
	var processedEvents []gin.H

	for _, event := range events {
		processedEvent := gin.H{
			"id":          event.ID,
			"description": event.Description,
			"year":        event.Year,
			"month":       event.Month,
			"day":         event.Day,
		}
		processedEvents = append(processedEvents, processedEvent)
	}

	return gin.H{"events": processedEvents}
}
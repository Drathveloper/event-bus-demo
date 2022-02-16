package event

import (
	"event-bus-demo/infrastructure/event_sourcing"
	"go.uber.org/zap"
)

func HandleError(logger *zap.Logger, event event_sourcing.Event, err error) event_sourcing.EventResult {
	if err != nil {
		logger.Error("error during event execution", zap.String("name", event.GetName()),
			zap.String("error", err.Error()))
		return event_sourcing.EventResult{
			Event: event,
		}
	}
	return event_sourcing.EventResult{
		Succeeded: true,
		Event:     event,
	}
}

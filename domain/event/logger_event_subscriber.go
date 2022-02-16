package event

import (
	"event-bus-demo/infrastructure/event_sourcing"
	"go.uber.org/zap"
)

type eventLoggerSubscriber struct {
	logger *zap.Logger
}

func NewEventLoggerSubscriber(logger *zap.Logger) event_sourcing.EventSubscriber {
	return &eventLoggerSubscriber{
		logger: logger,
	}
}

func (subscriber *eventLoggerSubscriber) Notify(result event_sourcing.EventResult) {
	subscriber.logger.Info("Subscriber received result", zap.Any("result", result))
}

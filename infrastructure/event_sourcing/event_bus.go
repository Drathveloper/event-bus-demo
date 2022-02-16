package event_sourcing

import (
	"context"
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
)

type EventBusChannel chan Event
type QuitSignalChannel chan bool

func NewBufferedEventChannel(bufferSize int) EventBusChannel {
	return make(chan Event, bufferSize)
}

func newQuitSignalChannel() QuitSignalChannel {
	return make(chan bool)
}

type EventBus interface {
	Run()
	Stop()
	Publish(event Event)
	RegisterHandler(topic string, eventBusHandler EventHandler)
	RegisterSubscriber(topic string, eventSubscriber EventSubscriber)
	UnregisterHandler(topic string, publisher EventHandler)
	UnregisterSubscriber(topic string, subscriber EventSubscriber)
}

type eventBus struct {
	logger             *zap.Logger
	handlerRegistry    map[string][]EventHandler
	subscriberRegistry map[string][]EventSubscriber
	semaphore          *semaphore.Weighted
	eventBusChannel    EventBusChannel
	quitSignalChannel  QuitSignalChannel
}

func NewEventBus(event EventBusChannel, maxWorkers int, logger *zap.Logger) EventBus {
	return &eventBus{
		eventBusChannel:    event,
		quitSignalChannel:  newQuitSignalChannel(),
		handlerRegistry:    make(map[string][]EventHandler),
		subscriberRegistry: make(map[string][]EventSubscriber),
		semaphore:          semaphore.NewWeighted(int64(maxWorkers)),
		logger:             logger,
	}
}

func (bus *eventBus) RegisterHandler(topic string, handler EventHandler) {
	bus.handlerRegistry[topic] = append(bus.handlerRegistry[topic], handler)
}

func (bus *eventBus) RegisterSubscriber(topic string, subscriber EventSubscriber) {
	bus.subscriberRegistry[topic] = append(bus.subscriberRegistry[topic], subscriber)
}

func (bus *eventBus) UnregisterHandler(topic string, handler EventHandler) {
	temp := bus.handlerRegistry[topic][:0]
	for _, element := range bus.handlerRegistry[topic] {
		if element != handler {
			temp = append(temp, element)
		}
	}
	bus.handlerRegistry[topic] = temp
}

func (bus *eventBus) UnregisterSubscriber(topic string, subscriber EventSubscriber) {
	temp := bus.subscriberRegistry[topic][:0]
	for _, element := range bus.subscriberRegistry[topic] {
		if element != subscriber {
			temp = append(temp, element)
		}
	}
	bus.subscriberRegistry[topic] = temp
}

func (bus *eventBus) Publish(event Event) {
	select {
	case bus.eventBusChannel <- event:
	default:
	}
}

func (bus *eventBus) Run() {
	go func() {
		for {
			select {
			case event := <-bus.eventBusChannel:
				bus.logger.Debug("new event received", zap.Any("event", event))
				_ = bus.semaphore.Acquire(context.Background(), 1)
				go func(event Event) {
					defer bus.semaphore.Release(1)
					bus.handleEvent(event)
				}(event)
			case <-bus.quitSignalChannel:
				bus.logger.Info("event bus signaled to stop")
				return
			default:
			}
		}
	}()
}

func (bus *eventBus) Stop() {
	go func() {
		bus.quitSignalChannel <- true
	}()
}

func (bus *eventBus) handleEvent(event Event) {
	eventTopic := event.GetTopic()
	foundHandlers := bus.handlerRegistry[eventTopic]
	if foundHandlers == nil {
		bus.logger.Debug("no bus handlers found for given event topic")
	} else {
		for _, handler := range foundHandlers {
			result := handler.Handle(event)
			bus.notifySubscribers(eventTopic, result)
		}
	}
}

func (bus *eventBus) notifySubscribers(topic string, result EventResult) {
	foundSubscribers := bus.subscriberRegistry[topic]
	if foundSubscribers == nil {
		bus.logger.Info("no bus subscribers found for given event topic")
	} else {
		for _, subscriber := range foundSubscribers {
			subscriber.Notify(result)
		}
	}
}

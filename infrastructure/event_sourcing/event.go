package event_sourcing

type Event interface {
	GetTopic() string
	GetName() string
}

type EventResult struct {
	Succeeded bool
	Event     Event
	Response  interface{}
}

type EventHandler interface {
	Handle(event Event) EventResult
}

type EventSubscriber interface {
	Notify(result EventResult)
}

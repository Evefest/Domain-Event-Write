package stream

import "github.com/Evefest/domainEventWrite/models"

type EventStream interface {
	SendCreatedEvent(event models.Event) error
}

var eventStream EventStream

func SetEventStream(stream EventStream) {
	eventStream = stream
}

func SendCreatedEvent(event models.Event) error {
	return eventStream.SendCreatedEvent(event)
}

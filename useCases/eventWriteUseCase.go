package useCases

import (
	"fmt"
	"github.com/Evefest/domainEventWrite/constants"
	"github.com/Evefest/domainEventWrite/exceptions"
	"github.com/Evefest/domainEventWrite/models"
	"github.com/Evefest/domainEventWrite/repositories"
	"github.com/Evefest/domainEventWrite/stream"
	"log"
)

func CreateEvent(event models.Event) (models.Event, exceptions.Response) {
	event.SetOnCreated()
	createdEvent, err := repositories.Create(event)
	if err != nil {
		log.Printf(constants.ErrorCreatingEvent, err.Error())
		return models.Event{}, exceptions.Response{}.Create(500, constants.InternalServerError)
	}
	go streamEventCreated(event)
	return createdEvent, exceptions.Response{}
}

func UpdateEvent(eventId string, event models.Event) (models.Event, exceptions.Response) {
	eventFound, errFindingEvent := findEventById(eventId)
	if errFindingEvent.ErrorCode != 0 {
		return models.Event{}, errFindingEvent
	}
	eventFound.SetOnUpdated(event)
	eventUpdated, err := repositories.Update(eventFound)
	if err != nil {
		log.Printf(constants.ErrorUpdatingEvent, err.Error())
		return models.Event{}, exceptions.Response{}.Create(500, constants.InternalServerError)
	}
	return eventUpdated, exceptions.Response{}
}

func ChangeEventVisibility(eventId string) exceptions.Response {
	eventFound, errFindingEvent := findEventById(eventId)
	if errFindingEvent.ErrorCode != 0 {
		return errFindingEvent
	}
	eventFound.ChangeVisibility()
	_, err := repositories.Update(eventFound)
	if err != nil {
		log.Printf(constants.ErrorUpdatingEvent, err.Error())
		return exceptions.Response{}.Create(500, constants.InternalServerError)
	}
	return exceptions.Response{}
}

func AddSubscribedUser(userId string, eventId string) {
	eventFound, errFindingEvent := findEventById(eventId)
	if errFindingEvent.ErrorCode != 0 {
		return
	}
	if eventFound.SubscribeUser(userId) {
		_, err := repositories.Update(eventFound)
		if err != nil {
			log.Printf(constants.ErrorUpdatingEvent, err.Error())
		}
	}
}

func RemoveSubscribedUser(userId string, eventId string) {
	eventFound, errFindingEvent := findEventById(eventId)
	if errFindingEvent.ErrorCode != 0 {
		return
	}
	if eventFound.UnsubscribeUser(userId) {
		_, err := repositories.Update(eventFound)
		if err != nil {
			log.Printf(constants.ErrorUpdatingEvent, err.Error())
		}
	}
}

func streamEventCreated(event models.Event) {
	if err := stream.SendCreatedEvent(event); err != nil {
		log.Printf(constants.ErrorSendingCreatedEvent, err.Error())
	}
}

func findEventById(eventId string) (models.Event, exceptions.Response) {
	eventFound, err := repositories.FindById(eventId)
	if err != nil {
		log.Printf(constants.ErrorFindingEvent, eventId, err.Error())
		return models.Event{}, exceptions.Response{}.Create(400, fmt.Sprintf(constants.EventNotFound, eventId))
	}
	return eventFound, exceptions.Response{}
}

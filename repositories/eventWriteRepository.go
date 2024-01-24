package repositories

import (
	"github.com/Evefest/domainEventWrite/models"
)

type EventRepository interface {
	Create(event models.Event) (models.Event, error)
	FindById(eventId string) (models.Event, error)
	Update(event models.Event) (models.Event, error)
	ExistsCategoryById(categoryId string) bool
	ExistsLocationById(locationId string) bool
}

var eventRepository EventRepository

func SetEventRepository(repository EventRepository) {
	eventRepository = repository
}

func Create(event models.Event) (models.Event, error) {
	return eventRepository.Create(event)
}

func FindById(eventId string) (models.Event, error) {
	return eventRepository.FindById(eventId)
}

func Update(event models.Event) (models.Event, error) {
	return eventRepository.Update(event)
}

func ExistsCategoryById(categoryId string) bool {
	return eventRepository.ExistsCategoryById(categoryId)
}

func ExistsLocationById(locationId string) bool {
	return eventRepository.ExistsLocationById(locationId)
}

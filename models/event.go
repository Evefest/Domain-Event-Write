package models

import (
	"github.com/Evefest/domainEventWrite/constants"
	"github.com/Evefest/domainEventWrite/utils"
	"log"
	"time"
)

type Event struct {
	Id              string    `json:"Id"`
	Name            string    `json:"Name"`
	StartTime       time.Time `json:"StartTime"`
	EndTime         time.Time `json:"EndTime"`
	Timezone        string    `json:"Timezone"`
	Details         string    `json:"Details"`
	EntryType       string    `json:"EntryType"`
	EventType       string    `json:"EventType"`
	Cost            float32   `json:"Cost"`
	AimedAt         string    `json:"AimedAt"`
	IsPublic        bool      `json:"IsPublic"`
	ThumbnailImage  string    `json:"ThumbnailImage"`
	Images          []string  `json:"Images"`
	LocationID      string    `json:"LocationId"`
	CategoryID      string    `json:"CategoryID"`
	SubscribedUsers []string  `json:"SubscribedUsers"`
	CreatedAt       time.Time `json:"CreatedAt"`
	UpdateAt        time.Time `json:"UpdateAt"`
	DisabledAt      time.Time `json:"DisabledAt"`
}

func (event *Event) SetOnCreated() {
	event.IsPublic = true
	event.CreatedAt = time.Now()
	event.Images = []string{}
	event.SubscribedUsers = []string{}
	event.setUpdatedDate()
}

func (event *Event) SetOnUpdated(eventUpdated Event) {
	event.Name = eventUpdated.Name
	event.StartTime = eventUpdated.StartTime
	event.EndTime = eventUpdated.EndTime
	event.Timezone = eventUpdated.Timezone
	event.Details = eventUpdated.Details
	event.EntryType = eventUpdated.EventType
	event.EventType = eventUpdated.EventType
	event.Cost = eventUpdated.Cost
	event.AimedAt = eventUpdated.AimedAt
	event.IsPublic = eventUpdated.IsPublic
	event.LocationID = eventUpdated.LocationID
	event.CategoryID = eventUpdated.CategoryID
	event.setUpdatedDate()
}

func (event *Event) DisableEvent() {
	event.IsPublic = false
	event.DisabledAt = time.Now()
}

func (event *Event) SubscribeUser(userId string) bool {
	if utils.ElementInSlice(userId, event.SubscribedUsers) {
		log.Printf(constants.UserAlreadySubscribed, userId)
		return false
	}
	event.SubscribedUsers = append(event.SubscribedUsers, userId)
	return true
}

func (event *Event) UnsubscribeUser(userId string) bool {
	if utils.ElementInSlice(userId, event.SubscribedUsers) {
		event.SubscribedUsers = utils.RemoveElementFromSlice(userId, event.SubscribedUsers)
		return true
	}
	log.Printf(constants.UserNotSubscribed, userId)
	return false
}

func (event *Event) setUpdatedDate() {
	event.UpdateAt = time.Now()
}

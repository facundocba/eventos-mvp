// file: application/service/event_service.go
package service

import (
	"github.com/prueba-hetmo/event-mvp/domain/model"
	"github.com/prueba-hetmo/event-mvp/domain/repository" // Asegúrate de que esta línea esté presente
)

type EventService interface {
	CreateEvent(event *model.Event) error
	UpdateEvent(event *model.Event) error
	DeleteEvent(id uint) error
	GetEventByID(id uint) (*model.Event, error)
	GetEvents(filter map[string]interface{}) ([]model.Event, error)
	EnrollUser(eventID uint, userID uint) error
	GetUserEvents(userID uint) ([]model.Event, error)
}

func NewEventService(repo repository.EventRepository) EventService {
	return &EventServiceImpl{EventRepo: repo}
}

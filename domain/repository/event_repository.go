// file: domain/repository/event_repository.go
package repository

import "github.com/prueba-hetmo/event-mvp/domain/model"

type EventRepository interface {
	Create(event *model.Event) error
	Update(event *model.Event) error
	Delete(id uint) error
	GetByID(id uint) (*model.Event, error)
	GetAll(filter map[string]interface{}) ([]model.Event, error)
	EnrollUser(eventID uint, userID uint) error
	GetUserEvents(userID uint) ([]model.Event, error)
}

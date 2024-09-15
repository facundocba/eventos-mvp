package service

import (
	"github.com/prueba-hetmo/event-mvp/domain/model"
	"github.com/prueba-hetmo/event-mvp/domain/repository"
)

type EventServiceImpl struct {
	EventRepo repository.EventRepository
}

func (s *EventServiceImpl) CreateEvent(event *model.Event) error {
	return s.EventRepo.Create(event)
}

func (s *EventServiceImpl) UpdateEvent(event *model.Event) error {
	return s.EventRepo.Update(event)
}

func (s *EventServiceImpl) DeleteEvent(id uint) error {
	return s.EventRepo.Delete(id)
}

func (s *EventServiceImpl) GetEventByID(id uint) (*model.Event, error) {
	return s.EventRepo.GetByID(id)
}

func (s *EventServiceImpl) GetEvents(filter map[string]interface{}) ([]model.Event, error) {
	return s.EventRepo.GetAll(filter)
}

func (s *EventServiceImpl) EnrollUser(eventID uint, userID uint) error {
	return s.EventRepo.EnrollUser(eventID, userID)
}

func (s *EventServiceImpl) GetUserEvents(userID uint) ([]model.Event, error) {
	return s.EventRepo.GetUserEvents(userID)
}

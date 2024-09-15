package persistence

import (
	"github.com/prueba-hetmo/event-mvp/domain/model"
	"github.com/prueba-hetmo/event-mvp/domain/repository"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) repository.EventRepository {
	return &EventRepositoryImpl{DB: db}
}

func (r *EventRepositoryImpl) Create(event *model.Event) error {
	return r.DB.Create(event).Error
}

func (r *EventRepositoryImpl) Update(event *model.Event) error {
	return r.DB.Save(event).Error
}

func (r *EventRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&model.Event{}, id).Error
}

func (r *EventRepositoryImpl) GetByID(id uint) (*model.Event, error) {
	var event model.Event
	if err := r.DB.First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *EventRepositoryImpl) GetAll(filter map[string]interface{}) ([]model.Event, error) {
	var events []model.Event
	query := r.DB
	for key, value := range filter {
		query = query.Where(key, value)
	}
	if err := query.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepositoryImpl) EnrollUser(eventID uint, userID uint) error {
	// Verifica si el usuario ya está inscrito
	var count int64
	if err := r.DB.Model(&model.EventUser{}).Where("event_id = ? AND user_id = ?", eventID, userID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil // Ya está inscrito
	}

	// Inscribe al usuario
	eventUser := &model.EventUser{EventID: eventID, UserID: userID}
	return r.DB.Create(eventUser).Error
}

func (r *EventRepositoryImpl) GetUserEvents(userID uint) ([]model.Event, error) {
	var eventIDs []uint
	if err := r.DB.Model(&model.EventUser{}).Where("user_id = ?", userID).Pluck("event_id", &eventIDs).Error; err != nil {
		return nil, err
	}

	var events []model.Event
	if err := r.DB.Where("id IN ?", eventIDs).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

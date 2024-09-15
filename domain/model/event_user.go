// file: domain/model/event_user.go

package model

type EventUser struct {
	ID      uint `gorm:"primaryKey"`
	EventID uint
	UserID  uint
}

// Add this line to Event model if it isn't already there
func (Event) TableName() string {
	return "events"
}

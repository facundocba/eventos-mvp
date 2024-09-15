package model

import "time"

type Event struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	Title            string    `json:"title"`
	DescriptionShort string    `json:"description_short"`
	DescriptionLong  string    `json:"description_long"`
	DateTime         time.Time `json:"date_time"`
	Organizer        string    `json:"organizer"`
	Location         string    `json:"location"`
	Status           string    `json:"status"`
	Participants     []*User   `gorm:"many2many:event_users;"` // Define la relación muchos a muchos
}

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	// Otros campos de usuario
}

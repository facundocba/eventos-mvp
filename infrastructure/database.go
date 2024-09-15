package infrastructure

import (
	"log"

	"github.com/prueba-hetmo/event-mvp/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/hetmodb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// AutoMigrate para el modelo Event
	err = db.AutoMigrate(&model.Event{}, &model.EventUser{})
	if err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	return db
}

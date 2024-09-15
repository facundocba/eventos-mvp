package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prueba-hetmo/event-mvp/api/handler"
	"github.com/prueba-hetmo/event-mvp/application/service"
	"github.com/prueba-hetmo/event-mvp/infrastructure"
	"github.com/prueba-hetmo/event-mvp/infrastructure/persistence"
	"github.com/prueba-hetmo/event-mvp/middleware"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Configura la base de datos
	db := infrastructure.SetupDatabase()

	// Configura los repositorios y servicios
	eventRepo := persistence.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	// Rutas públicas
	r.HandleFunc("/events", eventHandler.ListEvents).Methods("GET")
	r.HandleFunc("/events/{id:[0-9]+}", eventHandler.GetEvent).Methods("GET")
	r.HandleFunc("/user/events", eventHandler.GetUserEvents).Methods("GET")

	// Rutas protegidas
	r.Handle("/events", middleware.AdminMiddleware(http.HandlerFunc(eventHandler.CreateEvent))).Methods("POST")
	r.Handle("/events/{id:[0-9]+}", middleware.AdminMiddleware(http.HandlerFunc(eventHandler.UpdateEvent))).Methods("PUT")
	r.Handle("/events/{id:[0-9]+}", middleware.AdminMiddleware(http.HandlerFunc(eventHandler.DeleteEvent))).Methods("DELETE")
	r.Handle("/events/enroll", middleware.AuthMiddleware(http.HandlerFunc(eventHandler.EnrollUser))).Methods("POST")

	return r
}

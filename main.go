package main

import (
	"log"
	"net/http"

	"github.com/prueba-hetmo/event-mvp/api"
	"github.com/prueba-hetmo/event-mvp/infrastructure"
)

func main() {
	// Inicializa la base de datos
	db := infrastructure.SetupDatabase()
	if db == nil {
		log.Fatal("Failed to connect to the database.")
	}

	// Crea el router con la configuración adecuada
	router := api.NewRouter()

	// Inicia el servidor HTTP
	log.Fatal(http.ListenAndServe(":8080", router))
}

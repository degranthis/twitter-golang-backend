package main

import (
	"log"

	"github.com/degranthis/twitter-golang-backend/bd"
	"github.com/degranthis/twitter-golang-backend/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}

	handlers.Manejadores()

}

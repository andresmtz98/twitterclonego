package main

import (
	"github.com/andresmtz98/twitterclonego/bd"
	"github.com/andresmtz98/twitterclonego/handlers"
	"log"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}

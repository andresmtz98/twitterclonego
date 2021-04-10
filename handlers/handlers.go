package handlers

import (
	"github.com/andresmtz98/twitterclonego/middlew"
	"github.com/andresmtz98/twitterclonego/routers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux" //Un potente enrutador HTTP y comparador de URL para construir servidores web Go
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	// path: es el endpoint a capturar
	// func: se ejecuta un middleware para hacer las respectivas validaciones y luego ejecutar la función
	// que se le asigne vía parámetro
	// Methods: los métodos HTTP a capturar
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

package routers

import (
	"encoding/json"
	"github.com/andresmtz98/twitterclonego/bd"
	"github.com/andresmtz98/twitterclonego/models"
	"net/http"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) // Body es tipo Stream, una vez que se lee se destruye

	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese email", http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro de usuario ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

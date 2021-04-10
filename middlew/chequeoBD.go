package middlew

import (
	"github.com/andresmtz98/twitterclonego/bd"
	"net/http"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}

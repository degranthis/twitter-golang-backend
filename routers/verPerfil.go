package routers

import (
	"encoding/json"
	"net/http"

	"github.com/degranthis/twitter-golang-backend/bd"
)

/*VerPrfil permite extrar los valores del perfil */
func VerPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro", 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(perfil)

}

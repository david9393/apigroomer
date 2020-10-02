package routers

import (
	"encoding/json"
	"net/http"

	"github.com/david9393/apigroomer/bd"
)

func RazasLista(w http.ResponseWriter, r *http.Request) {

	results, err := bd.ListaRazas()
	if err != nil {
		http.Error(w, "Error al consultar los groomer"+err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

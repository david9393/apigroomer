package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/david9393/apigroomer/middlew"
	"github.com/david9393/apigroomer/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores seteo puertos */
func Manejadores() {
	router := mux.NewRouter()
	/*clientes*/
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/registrosall", middlew.ChequeoBD(routers.RegistroLista)).Methods("POST")
	/*groomer*/
	router.HandleFunc("/groomersall", middlew.ChequeoBD(routers.GroomerLista)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

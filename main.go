package main

import (
	"log"

	"github.com/david9393/apigroomer/bd"
	"github.com/david9393/apigroomer/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}

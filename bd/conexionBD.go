package bd

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "groomerapp"
)

var PostgresCN = ConectarBD()

func ConectarBD() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf(err.Error())
		return db
	}
	err = db.Ping()
	if err != nil {
		log.Printf(err.Error())
		return db
	}
	log.Printf("Conexion exitosaa")
	return db
}

func ChequeoConnection() int {
	err := PostgresCN.Ping()
	if err != nil {
		return 0
	}
	return 1
}

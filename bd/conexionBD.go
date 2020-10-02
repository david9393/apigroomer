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
	//host     = "ec2-54-243-67-199.compute-1.amazonaws.com"
	//port     = 5432
	//user     = "qtrvvsjxosvyjl"
	//password = "3e861e1966590ebed9b267be2f18ed2c49ed59155426374e7945576441cb4d1e"
	//dbname   = "dcn8aj88hcvj73"
)

var PostgresCN = ConectarBD()

func ConectarBD() *sql.DB {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// "password=%s dbname=%s",
	// host, port, user, password, dbname)
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

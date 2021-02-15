package bd

import (
	"database/sql"
	"log"

	"github.com/david9393/apigroomer/models"
)

func TraerUusario(u models.Login) (bool, models.Usuario) {

	usuario := models.Usuario{}

	db := Pool()
	const exec = `SELECT id,nombre from t05_usuarios where nombre=$1 and password=$2 `

	err := db.QueryRow(exec, u.Username, u.Password).Scan(&usuario.Id, &usuario.Nombre)
	switch {
	case err == sql.ErrNoRows:
		return false, usuario
	case err != nil:
		log.Fatal(err)
		return false, usuario
	default:
		return true, usuario
	}

}

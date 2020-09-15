package bd

import (
	"github.com/david9393/apigroomer/models"
)

func ListaGroomers(u models.Groomer) (models.GroomerList, error) {

	var groomers models.GroomerList
	db := PostgresCN

	const sql = `SELECT * from usp_groomer_select($1)`
	args := RetornarParametros(u)

	rows, err := db.Query(sql, args...)
	if err != nil {
		return groomers, err
	}
	for rows.Next() {
		groomer := &models.Groomer{}
		err = rows.Scan(&groomer.Id, &groomer.Direccion, &groomer.Nombre,
			&groomer.Telefono, &groomer.Calificacion, &groomer.Foto, &groomer.Descripcion)
		if err != nil {
			return groomers, err
		}
		groomers.ListGroomers = append(groomers.ListGroomers, groomer)
	}

	return groomers, nil
}

func RetornarParametros(u models.Groomer) ([]interface{}) {
	args := []interface{}{}
	if u.Id > 0 {
		args = append(args, &u.Id)
	} else {
		args = append(args, nil)
	}
	return args

}

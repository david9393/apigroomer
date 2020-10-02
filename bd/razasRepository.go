package bd

import (
	"github.com/david9393/apigroomer/models"
)

func ListaRazas() (models.RazasList, error) {

	var razas models.RazasList
	db := PostgresCN

	const sql = `SELECT id, nombre FROM t04_razas`

	rows, err := db.Query(sql)
	if err != nil {
		return razas, err
	}
	for rows.Next() {
		raza := &models.Razas{}
		err = rows.Scan(&raza.Id, &raza.Nombre)
		if err != nil {
			return razas, err
		}
		razas.ListRazas = append(razas.ListRazas, raza)
	}

	return razas, nil
}

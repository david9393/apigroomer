package bd

import (
	"log"

	"github.com/david9393/apigroomer/models"
)

func TraerMarcas() (error, []models.Marcas) {

	listMarcas := []models.Marcas{}

	db := Pool()
	const exec = `SELECT id, nombre
	FROM t08_marcas `

	rows, err := db.Query(exec)
	if err != nil {
		log.Fatal(err)
		return err, listMarcas
	}
	defer rows.Close()
	for rows.Next() {
		marca := &models.Marcas{}
		err = rows.Scan(&marca.Id, &marca.Nombre)
		if err != nil {
			log.Fatal(err)
			return err, listMarcas
		}
		listMarcas = append(listMarcas, *marca)
	}

	return nil, listMarcas

}
func CrearMarca(m models.Marcas) error {

	db := Pool()
	const exec = `INSERT INTO t08_marcas(
		 nombre) VALUES ($1);`

	_, err := db.Exec(exec, m.Nombre)
	if err != nil {
		return err
	}
	return nil
}
func UpdateMarcas(m models.Marcas) error {

	db := Pool()
	const exec = `UPDATE t08_marcas
	set nombre=$2
	WHERE id=$1`

	_, err := db.Exec(exec, m.Id, m.Nombre)
	if err != nil {
		return err
	}
	return nil
}

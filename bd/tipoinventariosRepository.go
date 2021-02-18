package bd

import (
	"log"

	"github.com/david9393/apigroomer/models"
)

func TraerTipoInventarios() (error, []models.TipoInventario) {

	listTipoInventarios := []models.TipoInventario{}

	db := Pool()
	const exec = `SELECT id, nombre
	FROM t09_tipo_inventarios `

	rows, err := db.Query(exec)
	if err != nil {
		log.Fatal(err)
		return err, listTipoInventarios
	}
	defer rows.Close()
	for rows.Next() {
		tipo := &models.TipoInventario{}
		err = rows.Scan(&tipo.Id, &tipo.Nombre)
		if err != nil {
			log.Fatal(err)
			return err, listTipoInventarios
		}
		listTipoInventarios = append(listTipoInventarios, *tipo)
	}

	return nil, listTipoInventarios

}
func CrearTipoInventarios(m models.TipoInventario) error {

	db := Pool()
	const exec = `INSERT INTO t09_tipo_inventarios(
		 nombre) VALUES ($1);`

	_, err := db.Exec(exec, m.Nombre)
	if err != nil {
		return err
	}
	return nil
}
func UpdateTiposInventario(m models.TipoInventario) error {

	db := Pool()
	const exec = `UPDATE t09_tipo_inventarios
	set nombre=$2
	WHERE id=$1`

	_, err := db.Exec(exec, m.Id, m.Nombre)
	if err != nil {
		return err
	}
	return nil
}

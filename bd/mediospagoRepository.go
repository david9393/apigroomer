package bd

import (
	"log"

	"github.com/david9393/apigroomer/models"
)

func TraerMediosPago() (error, []models.MediosPago) {

	listMedios := []models.MediosPago{}

	db := Pool()
	const exec = `SELECT id, nombre
	FROM t17_medios_pago `

	rows, err := db.Query(exec)
	if err != nil {
		log.Fatal(err)
		return err, listMedios
	}
	defer rows.Close()
	for rows.Next() {
		medio := &models.MediosPago{}
		err = rows.Scan(&medio.Id, &medio.Nombre)
		if err != nil {
			log.Fatal(err)
			return err, listMedios
		}
		listMedios = append(listMedios, *medio)
	}

	return nil, listMedios

}
func CrearMedio(m models.MediosPago) error {

	db := Pool()
	const exec = `INSERT INTO t17_medios_pago(
		 nombre) VALUES ($1);`

	_, err := db.Exec(exec, m.Nombre)
	if err != nil {
		return err
	}
	return nil
}
func UpdateMedio(m models.MediosPago) error {

	db := Pool()
	const exec = `UPDATE t17_medios_pago
	set nombre=$2
	WHERE id=$1`

	_, err := db.Exec(exec, m.Id, m.Nombre)
	if err != nil {
		return err
	}
	return nil
}

package bd

import (
	"database/sql"

	"github.com/david9393/apigroomer/models"
)

func TraerEncabezado() ([]models.MenuEncabezado, error) {

	var encabezadoList []models.MenuEncabezado

	db := Pool()
	const exec = `SELECT id, titulo, url from t02_menu_encabezado `

	rows, err := db.Query(exec)
	if err != nil {
		return encabezadoList, err
	}
	defer rows.Close()
	for rows.Next() {
		encabezado := &models.MenuEncabezado{}
		urlNull := sql.NullString{}
		err = rows.Scan(&encabezado.Id, &encabezado.Title, &urlNull)
		if err != nil {
			return encabezadoList, err
		}
		encabezado.Url = urlNull.String
		encabezadoList = append(encabezadoList, *encabezado)
	}

	return encabezadoList, nil
}

func TraerSubEncabezado(idpadre int) ([]models.MenuSubEncabezado, error) {

	var encabezadosubList []models.MenuSubEncabezado
	db := Pool()
	const sql = `SELECT id, titulo, url,idpadre,icon from t03_menu_sub_encabezado where idpadre=$1  `

	rows, err := db.Query(sql, idpadre)
	if err != nil {
		return encabezadosubList, err
	}
	defer rows.Close()
	for rows.Next() {
		encabezadoSub := &models.MenuSubEncabezado{}
		err = rows.Scan(&encabezadoSub.Id, &encabezadoSub.Description, &encabezadoSub.Url, &encabezadoSub.IdPadre, &encabezadoSub.Icon)
		if err != nil {
			return encabezadosubList, err
		}
		encabezadosubList = append(encabezadosubList, *encabezadoSub)
	}

	return encabezadosubList, nil
}
func TraerOpciones(idpadre int) ([]models.MenuOpciones, error) {

	var opcionList []models.MenuOpciones
	db := Pool()
	const sql = `SELECT id, titulo, url,idpadre,icon from t04_menu_opciones where idpadre=$1  `

	rows, err := db.Query(sql, idpadre)
	if err != nil {
		return opcionList, err
	}
	defer rows.Close()
	for rows.Next() {
		opcion := &models.MenuOpciones{}
		err = rows.Scan(&opcion.Id, &opcion.Description, &opcion.Url, &opcion.IdPadre, &opcion.Icon)
		if err != nil {
			return opcionList, err
		}
		opcionList = append(opcionList, *opcion)
	}
	rows.Close()

	return opcionList, nil
}

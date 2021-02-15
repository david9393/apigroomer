package bd

import (
	"log"

	"github.com/david9393/apigroomer/models"
)

func TraerValoresCategoriasInventarios(c models.ValoresCategoriasInventario) (error, []models.ValoresCategoriasInventario) {

	listValoresCategorias := []models.ValoresCategoriasInventario{}

	db := Pool()
	const exec = `SELECT id, id_categoria, nombre
	FROM t07_valores_categorias WHERE id_categoria=$1  `

	rows, err := db.Query(exec, c.IdCategorias)
	if err != nil {
		log.Fatal(err)
		return err, listValoresCategorias
	}
	defer rows.Close()
	for rows.Next() {
		valorCategoria := &models.ValoresCategoriasInventario{}
		err = rows.Scan(&valorCategoria.Id, &valorCategoria.IdCategorias, &valorCategoria.Nombre)
		if err != nil {
			log.Fatal(err)
			return err, listValoresCategorias
		}
		listValoresCategorias = append(listValoresCategorias, *valorCategoria)
	}

	return nil, listValoresCategorias

}
func CrearValoresCategoriasInventarios(c models.ValoresCategoriasInventario) error {

	db := Pool()
	const exec = `INSERT INTO t07_valores_categorias(
		id_categoria, nombre)
		VALUES ($1,$2)`

	_, err := db.Exec(exec, c.IdCategorias, c.Nombre)
	if err != nil {
		return err
	}
	return nil
}
func UpdateValoresCategoriasInventarios(c models.ValoresCategoriasInventario) error {

	db := Pool()
	const exec = `UPDATE t07_valores_categorias
	SET id_categoria=$2,nombre=$3
	WHERE id=$1`

	_, err := db.Exec(exec, c.Id, c.IdCategorias, c.Nombre)
	if err != nil {
		return err
	}
	return nil
}

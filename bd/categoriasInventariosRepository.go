package bd

import (
	"log"

	"github.com/david9393/apigroomer/models"
)

func TraerCategoriasInventarios() (error, []models.CategoriasInventario) {

	listCategorias := []models.CategoriasInventario{}

	db := Pool()
	const exec = `SELECT id, nombre
	FROM t06_categoria_inventarios `

	rows, err := db.Query(exec)
	if err != nil {
		log.Fatal(err)
		return err, listCategorias
	}
	defer rows.Close()
	for rows.Next() {
		categoria := &models.CategoriasInventario{}
		err = rows.Scan(&categoria.Id, &categoria.Nombre)
		if err != nil {
			log.Fatal(err)
			return err, listCategorias
		}
		listCategorias = append(listCategorias, *categoria)
	}

	return nil, listCategorias

}
func CrearCategoriasInventarios(c models.CategoriasInventario) error {

	db := Pool()
	const exec = `INSERT INTO public.t06_categoria_inventarios(
		nombre) VALUES ($1)`

	_, err := db.Exec(exec, c.Nombre)
	if err != nil {
		return err
	}
	return nil
}
func UpdateCategoriasInventarios(c models.CategoriasInventario) error {

	db := Pool()
	const exec = `UPDATE public.t06_categoria_inventarios
	SET nombre=$2
	WHERE id=$1`

	_, err := db.Exec(exec, c.Id, c.Nombre)
	if err != nil {
		return err
	}
	return nil
}

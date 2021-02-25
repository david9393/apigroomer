package bd

import (
	"database/sql"
	"log"

	"github.com/david9393/apigroomer/models"
)

//Consulta producto por un id
func TraerVendedores(v models.Vendedores) (error, []models.Vendedores) {

	listVendedores := []models.Vendedores{}

	db := Pool()
	const exec = `SELECT id, apellidos, nombres, direccion, email, telefono
	FROM public.t12_vendedores where  ($1 =0 OR id = $1) ORDER BY nombres ASC `

	rows, err := db.Query(exec, v.Id)
	if err != nil {
		log.Fatal(err)
		return err, listVendedores
	}
	defer rows.Close()
	for rows.Next() {
		vendedor := &models.Vendedores{}
		direccionNull := sql.NullString{}
		emailNull := sql.NullString{}
		telefonoNull := sql.NullString{}

		err = rows.Scan(&vendedor.Id, &vendedor.Apellidos, &vendedor.Nombres,
			&direccionNull, &emailNull, &telefonoNull)
		if err != nil {
			log.Fatal(err)
			return err, listVendedores
		}
		vendedor.Direccion = direccionNull.String
		vendedor.Email = emailNull.String
		vendedor.Telefono = telefonoNull.String
		listVendedores = append(listVendedores, *vendedor)
	}
	return nil, listVendedores
}
func CrearVendedores(v models.Vendedores) error {
	db := Pool()
	const exec = `INSERT INTO public.t12_vendedores(
		apellidos, nombres, direccion, email, telefono)
		VALUES ($1,$2,$3,$4,$5)`
	args := parametrosVendedores(v)
	_, err := db.Exec(exec, args...)
	if err != nil {
		return err
	}
	return nil
}
func UpdateVendedores(v models.Vendedores) error {

	db := Pool()
	const exec = `UPDATE public.t12_vendedores
	SET apellidos=$2, nombres=$3, direccion=$4, email=$5, telefono=$6
	WHERE id=$1`
	args := parametrosVendedores(v)
	_, err := db.Exec(exec, args...)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func parametrosVendedores(v models.Vendedores) []interface{} {
	args := []interface{}{}
	if v.Id > 0 {
		args = append(args, &v.Id)
	}
	args = append(args, &v.Apellidos)
	args = append(args, &v.Nombres)
	args = append(args, &v.Direccion)
	args = append(args, &v.Email)
	args = append(args, &v.Telefono)

	return args
}

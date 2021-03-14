package bd

import (
	"database/sql"
	"log"

	"github.com/david9393/apigroomer/models"
)

//Consulta producto por un id
func TraerClientes(p models.Cliente) (error, []models.Cliente) {

	listClientes := []models.Cliente{}

	db := Pool()
	const exec = `SELECT apellidos, id, nit, nombres, direccion, email, telefono,ispredeterminado,idvendedor,tipodocumento
	FROM public.t01_clientes where  ($1 =0 OR id = $1) `

	rows, err := db.Query(exec, p.Id)
	if err != nil {
		log.Fatal(err)
		return err, listClientes
	}
	defer rows.Close()
	for rows.Next() {
		cliente := &models.Cliente{}
		telefonoNull := sql.NullString{}
		emailNull := sql.NullString{}
		err = rows.Scan(&cliente.Apellidos, &cliente.Id, &cliente.Nit, &cliente.Nombres,
			&cliente.Direccion, &emailNull, &telefonoNull, &cliente.IsPredeterminado,
			&cliente.IdVendedor, &cliente.TipoDocumento)
		if err != nil {
			log.Fatal(err)
			return err, listClientes
		}
		cliente.Email = emailNull.String
		cliente.Telefono = telefonoNull.String
		listClientes = append(listClientes, *cliente)
	}
	return nil, listClientes
}
func CrearCliente(c models.Cliente) error {
	db := Pool()
	const exec = `INSERT INTO public.t01_clientes(
		apellidos,nit, nombres, direccion, email, telefono, ispredeterminado,idvendedor,tipodocumento)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	args := parametrosClientes(c)
	_, err := db.Exec(exec, args...)
	if err != nil {
		return err
	}
	return nil
}
func UpdateClientes(c models.Cliente) error {

	db := Pool()
	const exec = `UPDATE public.t01_clientes
	SET apellidos=$2,nit=$3, nombres=$4, direccion=$5, email=$6, telefono=$7,ispredeterminado=$8,idvendedor=$9,
    tipodocumento=$10 WHERE id=$1`
	args := parametrosClientes(c)
	_, err := db.Exec(exec, args...)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func parametrosClientes(c models.Cliente) []interface{} {
	args := []interface{}{}
	if c.Id > 0 {
		args = append(args, &c.Id)
	}
	args = append(args, &c.Apellidos)
	args = append(args, &c.Nit)
	args = append(args, &c.Nombres)
	args = append(args, &c.Direccion)
	args = append(args, &c.Email)
	args = append(args, &c.Telefono)
	args = append(args, &c.IsPredeterminado)
	args = append(args, &c.IdVendedor)
	args = append(args, &c.TipoDocumento)
	return args
}

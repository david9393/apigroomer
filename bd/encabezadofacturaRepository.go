package bd

import (
	"database/sql"
	"log"

	"github.com/david9393/apigroomer/models"
)

//Consulta producto por un id
func TraerFacturaId(p models.EncabezadoFactura) (error, []models.EncabezadoFactura) {

	listFacturas := []models.EncabezadoFactura{}

	db := Pool()
	const exec = `SELECT id, fecha, fechap, valorbruto, valordescuento, valoriva, valorneto, idvendedor, idcliente
	FROM public.t13_encabezado_factura where  ($1 =0 OR id = $1) `

	rows, err := db.Query(exec, p.Id)
	if err != nil {
		log.Fatal(err)
		return err, listFacturas
	}
	defer rows.Close()
	for rows.Next() {
		factura := &models.EncabezadoFactura{}
		err = rows.Scan(&factura.Id, &factura.Fecha, &factura.Fechap, &factura.ValorBruto,
			&factura.ValorDescuento, &factura.ValorIva, &factura.ValorNeto, &factura.IdVendedor,
			&factura.IdCliente)
		if err != nil {
			log.Fatal(err)
			return err, listFacturas
		}
		listFacturas = append(listFacturas, *factura)
	}
	return nil, listFacturas
}
func CrearEncabezadoFactura(p models.EncabezadoFactura) (error, int) {
	id := 0
	db := Pool()
	const exec = `INSERT INTO public.t13_encabezado_factura(
		fecha, fechap, valorbruto, valordescuento, valoriva, valorneto, idvendedor, idcliente)
	   VALUES ($1,NOW(), $2, $3, $4, $5, $6, $7) returning id`
	args := parametrosEncabezadoFactura(p)
	err := db.QueryRow(exec, args...).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		return err, id
	case err != nil:
		log.Fatal(err)
		return err, id
	default:
		return err, id
	}
}
func UpdateEncabezadoFactura(p models.EncabezadoFactura) error {

	db := Pool()
	const exec = `UPDATE public.t13_encabezado_factura
	SET  fecha=$2, fechap=NOW(), valorbruto=$3, valordescuento=$4, valoriva=$5, valorneto=$6, idvendedor=$7, idcliente=$8
	WHERE ID=$1`
	args := parametrosEncabezadoFactura(p)
	_, err := db.Exec(exec, args...)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func parametrosEncabezadoFactura(p models.EncabezadoFactura) []interface{} {
	args := []interface{}{}
	if p.Id > 0 {
		args = append(args, &p.Id)
	}
	args = append(args, &p.Fecha)
	args = append(args, &p.ValorBruto)
	args = append(args, &p.ValorDescuento)
	args = append(args, &p.ValorIva)
	args = append(args, &p.ValorNeto)
	args = append(args, &p.IdVendedor)
	args = append(args, &p.IdCliente)
	return args
}

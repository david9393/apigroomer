package bd

import (
	"log"

	"github.com/david9393/apigroomer/models"
)

//Consulta producto por un id
func TraerMovimientoFacturaId(p int32) (error, []models.MovimientoFactura) {

	listMovimientoFacturas := []models.MovimientoFactura{}

	db := Pool()
	const exec = `SELECT id, idencabezado, idproducto, cantidad, valorbruto, valoriva, valordescuento, valortotal
	FROM public.t14_movimiento_factura where  ($1 =0 OR idencabezado = $1) `

	rows, err := db.Query(exec, p)
	if err != nil {
		log.Fatal(err)
		return err, listMovimientoFacturas
	}
	defer rows.Close()
	for rows.Next() {
		facturaMovimiento := &models.MovimientoFactura{}
		err = rows.Scan(&facturaMovimiento.Id, &facturaMovimiento.IdEncabezado, &facturaMovimiento.Cantidad,
			&facturaMovimiento.ValorBruto, &facturaMovimiento.ValorIva, &facturaMovimiento.ValorDescuento,
			&facturaMovimiento.ValorNeto)
		if err != nil {
			log.Fatal(err)
			return err, listMovimientoFacturas
		}
		listMovimientoFacturas = append(listMovimientoFacturas, *facturaMovimiento)
	}
	return nil, listMovimientoFacturas
}
func CrearMovimientoFactura(p models.MovimientoFactura) error {
	db := Pool()
	const exec = `INSERT INTO public.t14_movimiento_factura(
		idencabezado, idproducto, cantidad, valorbruto, valoriva, valordescuento, valortotal)
	   VALUES ($1, $2, $3, $4, $5, $6, $7)`
	args := parametrosMovimientoFactura(p)
	_, err := db.Exec(exec, args...)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func UpdateMovimientoFactura(p models.MovimientoFactura) error {

	db := Pool()
	const exec = `UPDATE public.t14_movimiento_factura
	SET  idencabezado=$2, idproducto=$3, cantidad=$4, valorbruto=$5, valoriva=$6, 
	valordescuento=$7, valortotal=$8 WHERE id=$1`
	args := parametrosMovimientoFactura(p)
	_, err := db.Exec(exec, args...)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func parametrosMovimientoFactura(p models.MovimientoFactura) []interface{} {
	args := []interface{}{}
	if p.Id > 0 {
		args = append(args, &p.Id)
	}
	args = append(args, &p.IdEncabezado)
	args = append(args, &p.IdProducto)
	args = append(args, &p.Cantidad)
	args = append(args, &p.ValorBruto)
	args = append(args, &p.ValorIva)
	args = append(args, &p.ValorDescuento)
	args = append(args, &p.ValorNeto)
	return args
}

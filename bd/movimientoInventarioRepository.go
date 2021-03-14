package bd

import (
	"log"

	"github.com/david9393/apigroomer/models"
)

func CrearMovimientoInventario(p models.MovimientoInventario) error {
	db := Pool()
	const exec = `INSERT INTO public.t16_movimiento_inventario(
		idencabezado, isentrada, saldoanterior, cantidad, saldonuevo, movimiento, valorunitario, valortotal,fecha)
	   VALUES ($1,$2,$3,$4,$5,$6,$7,$8,NOW())`
	args := parametrosMovimientoInventario(p)
	_, err := db.Exec(exec, args...)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func parametrosMovimientoInventario(p models.MovimientoInventario) []interface{} {
	args := []interface{}{}
	if p.Id > 0 {
		args = append(args, &p.Id)
	}
	args = append(args, &p.IdEncabezado)
	args = append(args, &p.IsEntrada)
	args = append(args, &p.SaldoAnterior)
	args = append(args, &p.Cantidad)
	args = append(args, &p.SaldoNuevo)
	args = append(args, &p.Movimiento)
	args = append(args, &p.ValorUnitario)
	args = append(args, &p.ValorTotal)

	return args
}

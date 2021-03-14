package bd

import (
	"database/sql"

	"github.com/david9393/apigroomer/models"
)

func UpdateEncabezadoInventario(v models.EncabezadoInventario) (error, models.MovimientoInventario) {
	movimientoInventario := models.MovimientoInventario{}
	db := Pool()
	const exec = `UPDATE public.t15_encabezado_inventario
	SET  saldo=saldo+$2,valortotal=valorunitario*(saldo+$2)
	WHERE idproducto=$1 returning id,saldo,valorunitario`
	err := db.QueryRow(exec, v.IdProducto, v.Cantidad).Scan(&movimientoInventario.IdEncabezado, &movimientoInventario.SaldoNuevo, &movimientoInventario.ValorUnitario)
	switch {
	case err == sql.ErrNoRows:
		return err, movimientoInventario
	case err != nil:
		return err, movimientoInventario
	default:
		movimientoInventario.IsEntrada = true
		movimientoInventario.Cantidad = v.Cantidad * -1
		movimientoInventario.SaldoAnterior = movimientoInventario.SaldoNuevo - v.Cantidad
		movimientoInventario.Movimiento = "Factura"
		movimientoInventario.ValorTotal = movimientoInventario.ValorUnitario * (v.Cantidad * -1)
		return err, movimientoInventario
	}
}

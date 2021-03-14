package models

import (
	"time"
)

type EncabezadoFactura struct {
	Id             int       `json:"id"`
	Fecha          time.Time `json:"fecha"`
	Fechap         time.Time `json:"fechap"`
	ValorBruto     float32   `json:"valorBruto"`
	ValorDescuento float32   `json:"valorDescuento"`
	ValorIva       float32   `json:"valorIva"`
	ValorNeto      float32   `json:"valorNeto"`
	IdVendedor     int32     `json:"idVendedor"`
	IdCliente      int32     `json:"idCliente"`
}
type AddFacturaIn struct {
	Encabezado EncabezadoFactura   `json:"encabezado"`
	Movimiento []MovimientoFactura `json:"movimiento"`
}

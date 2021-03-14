package models

type MovimientoFactura struct {
	Id             int     `json:"id"`
	IdEncabezado   int     `json:"idEncabezado"`
	IdProducto     int     `json:"idProducto"`
	Cantidad       float32 `json:"cantidad"`
	ValorBruto     float32 `json:"valorBruto"`
	ValorDescuento float32 `json:"valorDescuento"`
	ValorIva       float32 `json:"valorIva"`
	ValorNeto      float32 `json:"valorNeto"`
}

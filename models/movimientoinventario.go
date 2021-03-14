package models

type MovimientoInventario struct {
	Id            int     `json:"id"`
	IdEncabezado  int     `json:"idEncabezado"`
	IsEntrada     bool    `json:"isEntrada"`
	SaldoAnterior float32 `json:"saldoAnterior"`
	Cantidad      float32 `json:"cantidad"`
	SaldoNuevo    float32 `json:"saldoNuevo"`
	Movimiento    string  `json:"movimiento"`
	ValorUnitario float32 `json:"valorUnitario"`
	ValorTotal    float32 `json:"valorTotal"`
}

package models

type EncabezadoInventario struct {
	Id            int     `json:"id"`
	IdProducto    int     `json:"idProducto"`
	Saldo         float32 `json:"saldo"`
	Cantidad      float32 `json:"cantidad"`
	ValorUnitario float32 `json:"valorUnitario"`
	ValorTotal    float32 `json:"valorTotal"`
}

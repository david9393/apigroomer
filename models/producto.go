package models

type Producto struct {
	Id                 int                  `json:"id"`
	Nombre             string               `json:"nombre"`
	Codigo             string               `json:"codigo"`
	IdTipo             int                  `json:"idTipo"`
	IdMarca            int                  `json:"idMarca"`
	Costo              float32              `json:"costo"`
	Precio             float32              `json:"precio"`
	PrecioMinimo       float32              `json:"precioMinimo"`
	Iva                float32              `json:"iva"`
	Minimo             float32              `json:"minimo"`
	Maximo             float32              `json:"maximo"`
	Servicio           bool                 `json:"servicio"`
	ControlaInventario bool                 `json:"controlaInventario"`
	Saldo              float32              `json:"saldo"`
	ListCategorias     []CategoriasProducto `json:"listCategorias"`
}

type CategoriasProducto struct {
	Id               int    `json:"id"`
	IdCategorias     int    `json:"idCategorias"`
	IdValorCategoria int    `json:"idValorCategoria"`
	IdProducto       int    `json:"idProducto"`
	Nombre           string `json:"nombre"`
}

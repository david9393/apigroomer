package models

type Cliente struct {
	Id        int    `json:"id"`
	Direccion string `json:"direccion"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Nit       string `json:"nit"`
}

type ClientesList struct {
	ListClientes []*Cliente `json:"listClientes"`
}
type ClienteIn struct {
	Cliente *Cliente `json:"cliente"`
	IdPadre int      `json:"idPadre"`
}

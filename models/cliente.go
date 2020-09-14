package models

type Cliente struct {
	Id        int    `json:"id"`
	Direccion string `json:"direccion"`
	Nombre    string `json:"nombre"`
	Telefono  int    `json:"telefono"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type ClientesList struct {
	ListClientes []*Cliente `json:"listClientes"`
}

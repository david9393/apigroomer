package models

type Vendedores struct {
	Id        int    `json:"id"`
	Apellidos string `json:"apellidos"`
	Nombres   string `json:"nombres"`
	Direccion string `json:"direccion"`
	Email     string `json:"email"`
	Telefono  string `json:"telefono"`
}

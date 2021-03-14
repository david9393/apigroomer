package models

type Cliente struct {
	Id               int    `json:"id"`
	Apellidos        string `json:"apellidos"`
	Nit              string `json:"nit"`
	Nombres          string `json:"nombres"`
	Direccion        string `json:"direccion"`
	Email            string `json:"email"`
	Telefono         string `json:"telefono"`
	IsPredeterminado bool   `json:"isPredeterminado"`
	IdVendedor       int    `json:"idVendedor"`
	TipoDocumento    string `json:"tipoDocumento"`
}

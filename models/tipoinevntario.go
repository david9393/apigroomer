package models

type TipoInventario struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}
type ListTipoInventarios struct {
	ListTipoInventarios []TipoInventario `json:"listTipoInventarios"`
}

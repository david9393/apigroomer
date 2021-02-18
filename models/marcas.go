package models

type Marcas struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}
type ListMarcas struct {
	ListMarcas []Marcas `json:"listMarcas"`
}

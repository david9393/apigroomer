package models

type ValoresCategoriasInventario struct {
	Id           int    `json:"id"`
	IdCategorias int    `json:"idCategorias"`
	Nombre       string `json:"nombre"`
}
type ListValoresCategoriasInventarios struct {
	ListValoresCategoriasInventarios []ValoresCategoriasInventario `json:"listValoresCategoriasInventarios"`
}

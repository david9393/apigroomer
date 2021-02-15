package models

type CategoriasInventario struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`

}
type ListCategoriasInventarios struct {
	ListCategoriasInventarios []CategoriasInventario `json:"listCategoriasInventarios"`
}

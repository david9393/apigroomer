package models

type Usuario struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Password string `json:"password"`
}

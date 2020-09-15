package models

type Groomer struct {
	Id               int    `json:"id"`
	Direccion        string `json:"direccion"`
	Nombre           string `json:"nombre"`
	Telefono         int    `json:"telefono"`
	Calificacion     int    `json:"calificacion"`
	Foto             string `json:"foto"`
	Descripcion      string `json:"descripcion"`
}

type GroomerList struct {
	ListGroomers []*Groomer `json:"listGroomers"`
}

package models

type Servicios struct {
	Id         int    `json:"id"`
	Idraza     int    `json:"idraza"`
	Iscachorro bool   `json:"iscachorro"`
	Nombre     string `json:"nombre"`
	Detalle    string `json:"detalle"`
	Foto       string `json:"foto"`
	Precio     int    `json:"precio"`
}

type ListServicios struct {
	ListServicios []*Servicios `json:"listServicios"`
}

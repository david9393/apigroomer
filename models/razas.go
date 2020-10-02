package models

type Razas struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

type RazasList struct {
	ListRazas []*Razas `json:"listRazas"`
}

package models

type MenuOut struct {
	ListMenu []MenuEncabezado `json:"listMenu"`
}

type MenuEncabezado struct {
	Id         int                 `json:"id"`
	Title      string              `json:"title"`
	Url        string              `json:"url"`
	ListOption []MenuSubEncabezado `json:"listOption"`
}
type MenuSubEncabezado struct {
	Id          int            `json:"id"`
	Description string         `json:"description"`
	Url         string         `json:"url"`
	IdPadre     string         `json:"idPadre"`
	Icon        string         `json:"icon"`
	ListOption  []MenuOpciones `json:"listOption"`
}

type MenuOpciones struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Url         string `json:"url"`
	IdPadre     string `json:"idPadre"`
	Icon        string `json:"icon"`
}

package routers

import (
	"net/http"

	"github.com/david9393/apigroomer/bd"
	"github.com/david9393/apigroomer/models"
	"github.com/labstack/echo"
)

func TraerMenu(c echo.Context) error {
	var encabezado []models.MenuEncabezado
	var menu models.MenuOut

	encabezado, err := bd.TraerEncabezado()
	menu.ListMenu = encabezado
	if err != nil {
		response := NewResponse(Error, "Hubo un problema al cargar el menu", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	for i := len(menu.ListMenu) - 1; i >= 0; i-- {
		res, err := bd.TraerSubEncabezado(menu.ListMenu[i].Id)
		if err != nil {
			response := NewResponse(Error, "Hubo un problema al cargar el menu", nil)
			return c.JSON(http.StatusInternalServerError, response)
		} else {
			for i := len(res) - 1; i >= 0; i-- {
				res1, err1 := bd.TraerOpciones(res[i].Id)
				if err1 != nil {
					response := NewResponse(Error, "Hubo un problema al cargar el menu", nil)
					return c.JSON(http.StatusInternalServerError, response)
				}
				res[i].ListOption = res1
			}
			menu.ListMenu[i].ListOption = res
		}
	}
	return c.JSON(http.StatusOK, menu)

}

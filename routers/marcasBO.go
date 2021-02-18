package routers

import (
	"net/http"

	"github.com/david9393/apigroomer/bd"
	"github.com/david9393/apigroomer/models"
	"github.com/labstack/echo"
)

func CrearMarcas(c echo.Context) error {

	data := models.Marcas{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if data.Id > 0 {
		err = bd.UpdateMarcas(data)
	} else {
		err = bd.CrearMarca(data)
	}

	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusAccepted, resp)
	}
	resp := NewResponse("OK", "OK", nil)
	return c.JSON(http.StatusOK, resp)
}

func TraerMarcas(c echo.Context) error {
	err, listMarcas := bd.TraerMarcas()
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := NewResponse("OK", "OK", listMarcas)

	return c.JSON(http.StatusOK, resp)
}

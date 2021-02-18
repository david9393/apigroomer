package routers

import (
	"net/http"

	"github.com/david9393/apigroomer/bd"
	"github.com/david9393/apigroomer/models"
	"github.com/labstack/echo"
)

func CrearTipoInevntarios(c echo.Context) error {

	data := models.TipoInventario{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if data.Id > 0 {
		err = bd.UpdateTiposInventario(data)
	} else {
		err = bd.CrearTipoInventarios(data)
	}

	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusAccepted, resp)
	}
	resp := NewResponse("OK", "OK", nil)
	return c.JSON(http.StatusOK, resp)
}

func TraerTipoInventarios(c echo.Context) error {
	err, listTiposInventarios := bd.TraerTipoInventarios()
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := NewResponse("OK", "OK", listTiposInventarios)

	return c.JSON(http.StatusOK, resp)
}

package routers

import (
	"net/http"

	"github.com/david9393/apigroomer/bd"
	"github.com/david9393/apigroomer/models"
	"github.com/labstack/echo"
)

func CrearCliente(c echo.Context) error {

	data := models.Cliente{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if data.Id > 0 {
		err = bd.UpdateClientes(data)
	} else {
		err = bd.CrearCliente(data)
	}

	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusAccepted, resp)
	}
	resp := NewResponse("OK", "OK", nil)
	return c.JSON(http.StatusOK, resp)
}

func TraerCliente(c echo.Context) error {
	data := models.Cliente{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	err, listClientes := bd.TraerClientes(data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := NewResponse("OK", "OK", listClientes)

	return c.JSON(http.StatusOK, resp)
}

package routers

import (
	"net/http"
	"strings"

	"github.com/david9393/apigroomer/bd"
	"github.com/david9393/apigroomer/models"
	"github.com/labstack/echo"
)

func CrearVendedor(c echo.Context) error {

	data := models.Vendedores{}
	err := c.Bind(&data)
	data.Nombres = strings.Title(data.Nombres)
	data.Apellidos = strings.Title(data.Apellidos)
	data.Email = strings.ToLower(data.Email)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if data.Id > 0 {
		err = bd.UpdateVendedores(data)
	} else {
		err = bd.CrearVendedores(data)
	}

	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusAccepted, resp)
	}
	resp := NewResponse("OK", "OK", nil)
	return c.JSON(http.StatusOK, resp)
}

func TraerVendedores(c echo.Context) error {
	data := models.Vendedores{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	err, listVendedores := bd.TraerVendedores(data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := NewResponse("OK", "OK", listVendedores)

	return c.JSON(http.StatusOK, resp)
}

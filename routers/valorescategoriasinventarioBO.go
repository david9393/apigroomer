package routers

import (
	"net/http"

	"github.com/david9393/apigroomer/bd"
	"github.com/david9393/apigroomer/models"
	"github.com/labstack/echo"
)

func CrearValoresCategoriasInventarios(c echo.Context) error {

	data := []models.ValoresCategoriasInventario{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	for _, valor := range data {
		if valor.Id > 0 {
			err = bd.UpdateValoresCategoriasInventarios(valor)

		} else {
			err = bd.CrearValoresCategoriasInventarios(valor)

		}
		if err != nil {
			resp := NewResponse(Error, ErrorMessage, nil)
			return c.JSON(http.StatusBadRequest, resp)
		}
	}
	resp := NewResponse("OK", "OK", nil)
	return c.JSON(http.StatusOK, resp)
}

func TraerValoresCategoriasInventarios(c echo.Context) error {
	data := models.ValoresCategoriasInventario{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	err, listCatecorias := bd.TraerValoresCategoriasInventarios(data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := NewResponse("OK", "OK", listCatecorias)

	return c.JSON(http.StatusOK, resp)
}
func UpdateValoresCategoriasInventarios(c echo.Context) error {

	data := models.ValoresCategoriasInventario{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	err = bd.UpdateValoresCategoriasInventarios(data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := NewResponse("OK", "OK", nil)
	return c.JSON(http.StatusOK, resp)
}

package routers

import (
	"log"
	"net/http"

	"github.com/david9393/apigroomer/bd"
	"github.com/david9393/apigroomer/models"
	"github.com/labstack/echo"
)

func CrearProductos(c echo.Context) error {
	id := 0
	data := models.Producto{}
	err := c.Bind(&data)
	if err != nil {
		log.Fatal(err)
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if data.Id > 0 {
		err = bd.UpdateProductos(data)
		id = data.Id
	} else {
		err, id = bd.CrearProducto(data)
	}

	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusAccepted, resp)
	}
	for _, v := range data.ListCategorias {
		v.IdProducto = id
		if v.Id > 0 {
			err = bd.UpdateCategoriaProducto(v)
		} else {
			err = bd.CrearCategoriaProducto(v)
		}
		if err != nil {
			resp := NewResponse(Error, ErrorMessage, nil)
			return c.JSON(http.StatusAccepted, resp)
		}
	}
	resp := NewResponse("OK", "OK", nil)
	return c.JSON(http.StatusOK, resp)
}

func TraerProductoId(c echo.Context) error {
	data := models.Producto{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	err, listProductos := bd.TraerProductoId(data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	err, listProductos[0].ListCategorias = bd.TraerCategoriaProducto(data)
	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := NewResponse("OK", "OK", listProductos)

	return c.JSON(http.StatusOK, resp)
}

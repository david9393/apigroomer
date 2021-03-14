package routers

import (
	"log"
	"net/http"

	"github.com/david9393/apigroomer/bd"
	"github.com/david9393/apigroomer/models"
	"github.com/labstack/echo"
)

func CrearFactura(c echo.Context) error {
	id := 0
	movimientoInventario := models.MovimientoInventario{}
	data := models.AddFacturaIn{}
	encabezadoinv := models.EncabezadoInventario{}
	err := c.Bind(&data)
	if err != nil {
		log.Fatal(err)
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if data.Encabezado.Id > 0 {
		err = bd.UpdateEncabezadoFactura(data.Encabezado)
		id = data.Encabezado.Id
	} else {
		err, id = bd.CrearEncabezadoFactura(data.Encabezado)
	}

	if err != nil {
		resp := NewResponse(Error, ErrorMessage, nil)
		return c.JSON(http.StatusAccepted, resp)
	}
	for _, v := range data.Movimiento {
		v.IdEncabezado = id
		err = bd.CrearMovimientoFactura(v)
		if err != nil {
			resp := NewResponse(Error, ErrorMessage, nil)
			return c.JSON(http.StatusAccepted, resp)
		}
		encabezadoinv.IdProducto = v.IdProducto
		encabezadoinv.Cantidad = v.Cantidad * -1
		err, movimientoInventario = bd.UpdateEncabezadoInventario(encabezadoinv)
		if err != nil {
			resp := NewResponse(Error, ErrorMessage, nil)
			return c.JSON(http.StatusAccepted, resp)
		}
		err = bd.CrearMovimientoInventario(movimientoInventario)
		if err != nil {
			resp := NewResponse(Error, ErrorMessage, nil)
			return c.JSON(http.StatusAccepted, resp)
		}
	}
	resp := NewResponse("OK", "OK", nil)
	return c.JSON(http.StatusOK, resp)
}

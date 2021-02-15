package routers

import (
	"net/http"

	"github.com/david9393/apigroomer/authorization"
	"github.com/david9393/apigroomer/bd"
	"github.com/david9393/apigroomer/models"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	token := ""
	dataToken := map[string]string{"token": token}
	data := models.Login{}
	usuario := models.Usuario{}
	var isvalido bool
	err := c.Bind(&data)
	if err != nil {
		dataToken = map[string]string{"token": token}

		// resp := NewResponse(Error, "estructura no válida", nil)
		return c.JSON(http.StatusBadRequest, dataToken)
	}
	isvalido, usuario = bd.TraerUusario(data)
	if !isvalido {
		// resp := NewResponse(Error, "usuario o contraseña no válidos", nil)
		return c.JSON(http.StatusOK, dataToken)
	}

	token, err = authorization.GenerateToken(&usuario)
	if err != nil {
		// resp := NewResponse(Error, "no se pudo generar el token", nil)
		return c.JSON(http.StatusOK, token)
	}

	dataToken = map[string]string{"token": token}
	return c.JSON(http.StatusOK, dataToken)
}


package handlers

import (
	"github.com/david9393/apigroomer/routers"
	"github.com/labstack/echo"
)

/* Manejadores seteo puertos */
func RouteMenu(e *echo.Echo) {
	menu := e.Group("/menu")
	menu.POST("/all", routers.TraerMenu)
}
func RouteLogin(e *echo.Echo) {
	e.POST("/auth/login", routers.Login)
}

func RouteCategoriasInventario(e *echo.Echo) {
	menu := e.Group("/categoriasinventario")
	menu.POST("/add", routers.CrearCategoriasInventarios)
	menu.POST("/all", routers.TraerCategoriasInventarios)
	menu.POST("/update", routers.UpdateCategoriasInventarios)
	menu.POST("/addvalores", routers.CrearValoresCategoriasInventarios)
	menu.POST("/allvalores", routers.TraerValoresCategoriasInventarios)
	menu.POST("/updatevalores", routers.UpdateValoresCategoriasInventarios)
}

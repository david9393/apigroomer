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
func RouteMarcas(e *echo.Echo) {
	menu := e.Group("/marcas")
	menu.POST("/add", routers.CrearMarcas)
	menu.POST("/all", routers.TraerMarcas)

}
func RouteTiposInventario(e *echo.Echo) {
	menu := e.Group("/tipoinventarios")
	menu.POST("/add", routers.CrearTipoInevntarios)
	menu.POST("/all", routers.TraerTipoInventarios)

}

func RouteProductos(e *echo.Echo) {
	menu := e.Group("/productos")
	menu.POST("/add", routers.CrearProductos)
	menu.POST("/all", routers.TraerProductoId)

}

func RouteClientes(e *echo.Echo) {
	menu := e.Group("/clientes")
	menu.POST("/add", routers.CrearCliente)
	menu.POST("/all", routers.TraerCliente)

}
func RouteVendedores(e *echo.Echo) {
	menu := e.Group("/vendedores")
	menu.POST("/add", routers.CrearVendedor)
	menu.POST("/all", routers.TraerVendedores)

}

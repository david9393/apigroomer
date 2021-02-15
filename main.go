package main

import (
	"log"

	"github.com/david9393/apigroomer/authorization"
	"github.com/david9393/apigroomer/bd"
	"github.com/david9393/apigroomer/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}
	bd.ConectarBD()
	log.Println("Conexion exitosa")

	e := echo.New()
	e.Use(middleware.CORS())
	handlers.RouteMenu(e)
	handlers.RouteLogin(e)
	handlers.RouteCategoriasInventario(e)
	log.Println("Servidor iniciado en el puerto 8080")

	err = e.Start(":8080")
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}

}

// func GeneratePdf(filename string) error {

// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.AddPage()
// 	pdf.SetFont("Arial", "B", 16)

// 	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
// 	pdf.CellFormat(190, 7, "Welcome to golangcode.com", "0", 0, "CM", false, 0, "")

// 	// ImageOptions(src, x, y, width, height, flow, options, link, linkStr)

// 	return pdf.OutputFileAndClose(filename)
// }

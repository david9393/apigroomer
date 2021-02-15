package bd

import (
	"strconv"

	"github.com/david9393/apigroomer/models"
)

func InsertarCliente(u models.ClienteIn) (string, error) {

	db := Pool()
	sqlStatement := `INSERT INTO public.t01_clientes(
		apellidos, nit, nombres, direccion)
		VALUES ($1, $2, $3, $4);`
	args := RetornarParametrosCliente(*u.Cliente)

	_, err := db.Exec(sqlStatement, args...)
	if err != nil {
		return "Ocurrion un Error", err
	}
	db.Close()
	return "Se creo correctamente el cliente" + strconv.Itoa(u.IdPadre), nil
}
func RetornarParametrosCliente(u models.Cliente) []interface{} {
	args := []interface{}{}
	args = append(args, &u.Apellidos)
	args = append(args, &u.Nit)
	args = append(args, &u.Nombres)
	args = append(args, &u.Direccion)
	return args
}

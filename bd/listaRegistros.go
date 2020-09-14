package bd

import (
	"github.com/david9393/apigroomer/models"
)

func ListaRegistros(u models.Cliente) (models.ClientesList, error) {

	var registros models.ClientesList
	db := PostgresCN
	u.Password, _ = EncriptarPassword(u.Password)
	const sql = `select * from usp_clientes_select($1,$2,$3)`
	args := []interface{}{}
	if u.Id > 0 {
		args = append(args, &u.Id)
	} else {
		args = append(args, nil)
	}
	if u.Email != "" {
		args = append(args, &u.Email)
	} else {
		args = append(args, nil)
	}
	if u.Nombre != "" {
		args = append(args, &u.Nombre)
	} else {
		args = append(args, nil)
	}
	rows, err := db.Query(sql, args...)
	if err != nil {
		return registros, err
	}
	for rows.Next() {
		registro := &models.Cliente{}
		err = rows.Scan(&registro.Direccion, &registro.Id, &registro.Nombre,
			&registro.Telefono, &registro.Email)
		if err != nil {
			return registros, err
		}
		registros.ListClientes = append(registros.ListClientes, registro)
	}

	return registros, nil
}

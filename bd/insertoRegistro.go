package bd

import (
	"github.com/david9393/apigroomer/models"
)

/* Inserto registro*/
func InsertarRegistro(u models.Cliente) (int, error) {

	db := PostgresCN
	u.Password, _ = EncriptarPassword(u.Password)
	const sql = `select * from  usp_cliente_insert($1,$2,$3,$4,$5)`

	row := db.QueryRow(sql, u.Direccion, u.Email, u.Nombre, u.Password, u.Telefono)
	id := 0
	err := row.Scan(&id)
	if id == 0 {
		return id, err
	}
	return id, nil
}
func ChequeoExisteUsuario(email string) bool {

	db := PostgresCN
	const sql = `select * from  usp_cliente_email($1)`

	row := db.QueryRow(sql, email)

	id := 0
	row.Scan(&id)

	if id != 0 {
		return true
	}
	return false
}

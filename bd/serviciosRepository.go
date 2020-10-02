package bd

import (
	"github.com/david9393/apigroomer/models"
)

func ListaServicios(u models.Servicios) (models.ListServicios, error) {

	var servicios models.ListServicios
	db := PostgresCN

	const sql = `SELECT id,idraza, iscachorro, nombre, detalle, 
	case when foto isnull then '' else foto end, precio
	 FROM t06_servicios WHERE idraza=$1 AND iscachorro=$2;`

	rows, err := db.Query(sql, u.Idraza, u.Iscachorro)
	if err != nil {
		return servicios, err
	}
	for rows.Next() {
		servicio := &models.Servicios{}
		err = rows.Scan(&servicio.Id, &servicio.Idraza, &servicio.Iscachorro, &servicio.Nombre,
			&servicio.Detalle, &servicio.Foto, &servicio.Precio)
		if err != nil {
			return servicios, err
		}
		servicios.ListServicios = append(servicios.ListServicios, servicio)
	}

	return servicios, nil
}

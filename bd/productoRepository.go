package bd

import (
	"database/sql"
	"log"

	"github.com/david9393/apigroomer/models"
)

//Consulta producto por un id
func TraerProductoId(p models.Producto) (error, []models.Producto) {

	listProductos := []models.Producto{}

	db := Pool()
	const exec = `SELECT id, nombre, codigo, idtipo, idmarca, costo, precio, 
	preciominimo, iva, minimo, maximo, servicio, controlainventario
	FROM public.t10_productos where  ($1 =0 OR id = $1) `

	rows, err := db.Query(exec, p.Id)
	if err != nil {
		log.Fatal(err)
		return err, listProductos
	}
	defer rows.Close()
	for rows.Next() {
		producto := &models.Producto{}
		err = rows.Scan(&producto.Id, &producto.Nombre, &producto.Codigo, &producto.IdTipo,
			&producto.IdMarca, &producto.Costo, &producto.Precio, &producto.PrecioMinimo, &producto.Iva,
			&producto.Minimo, &producto.Maximo, &producto.Servicio, &producto.ControlaInventario)
		if err != nil {
			log.Fatal(err)
			return err, listProductos
		}
		listProductos = append(listProductos, *producto)
	}
	return nil, listProductos
}
func CrearProducto(p models.Producto) (error, int) {
	id := 0
	db := Pool()
	const exec = `INSERT INTO public.t10_productos(
		nombre, codigo, idtipo, idmarca, costo, precio, preciominimo, iva, minimo, maximo, servicio, controlainventario)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning id`
	args := parametrosProductos(p)
	err := db.QueryRow(exec, args...).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		return err, id
	case err != nil:
		log.Fatal(err)
		return err, id
	default:
		return err, id
	}
}
func UpdateProductos(p models.Producto) error {

	db := Pool()
	const exec = `UPDATE public.t10_productos
	SET  nombre=$2, codigo=$3, idtipo=$4, idmarca=$5, costo=$6, precio=$7, preciominimo=$8, 
	iva=$9, minimo=$10, maximo=$11, servicio=$12, controlainventario=$13
	WHERE id=$1`
	args := parametrosProductos(p)
	_, err := db.Exec(exec, args...)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func UpdateCategoriaProducto(p models.CategoriasProducto) error {
	db := Pool()
	const exec = `UPDATE public.t11_categorias_producto
	SET  idvalorcategoria=$2, idproducto=$3
	WHERE id=$1`
	_, err := db.Exec(exec, p.Id, p.IdValorCategoria, p.IdProducto)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func CrearCategoriaProducto(p models.CategoriasProducto) error {
	db := Pool()
	const exec = `INSERT INTO public.t11_categorias_producto(
		idvalorcategoria, idproducto)
		VALUES ($1,$2)`
	_, err := db.Exec(exec, p.IdValorCategoria, p.IdProducto)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func TraerCategoriaProducto(p models.Producto) (error, []models.CategoriasProducto) {
	listCategoriasProducto := []models.CategoriasProducto{}

	db := Pool()
	const exec = `SELECT t11_categorias_producto.id, 
	t11_categorias_producto.idvalorcategoria,
	t11_categorias_producto.idproducto,
	t06_categoria_inventarios.nombre
    FROM public.t11_categorias_producto
    INNER JOIN t07_valores_categorias
    ON t07_valores_categorias.id = idvalorcategoria
    INNER JOIN t06_categoria_inventarios
    ON t06_categoria_inventarios.id=t07_valores_categorias.id_categoria
    where idproducto=$1`

	rows, err := db.Query(exec, p.Id)
	if err != nil {
		log.Fatal(err)
		return err, listCategoriasProducto
	}
	defer rows.Close()
	for rows.Next() {
		categoria := &models.CategoriasProducto{}
		err = rows.Scan(&categoria.Id, &categoria.IdValorCategoria, &categoria.IdProducto, &categoria.Nombre)
		if err != nil {
			log.Fatal(err)
			return err, listCategoriasProducto
		}
		listCategoriasProducto = append(listCategoriasProducto, *categoria)
	}

	return nil, listCategoriasProducto
}
func parametrosProductos(p models.Producto) []interface{} {
	args := []interface{}{}
	if p.Id > 0 {
		args = append(args, &p.Id)
	}
	args = append(args, &p.Nombre)
	args = append(args, &p.Codigo)
	args = append(args, &p.IdTipo)
	args = append(args, &p.IdMarca)
	args = append(args, &p.Costo)
	args = append(args, &p.Precio)
	args = append(args, &p.PrecioMinimo)
	args = append(args, &p.Iva)
	args = append(args, &p.Minimo)
	args = append(args, &p.Maximo)
	args = append(args, &p.Servicio)
	args = append(args, &p.ControlaInventario)
	return args
}

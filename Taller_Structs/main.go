package main

import "fmt"

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}

func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	p.Precio = nuevoPrecio
}

func (p *Producto) ActualizarCantidad(nuevaCantidad int) {
	p.Cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Println("---------------------------")
	fmt.Println("ID:", p.ID)
	fmt.Println("Nombre:", p.Nombre)
	fmt.Printf("Precio: $%.2f\n", p.Precio)
	fmt.Println("Cantidad:", p.Cantidad)
	fmt.Println("---------------------------")
}

type Inventario struct {
	Productos map[int]Producto
}

func NuevoInventario() *Inventario {
	return &Inventario{
		Productos: make(map[int]Producto),
	}
}

func (i *Inventario) AgregarProducto(p Producto) {
	if _, existe := i.Productos[p.ID]; existe {
		fmt.Println("Ya existe un producto con ese ID")
		return
	}

	i.Productos[p.ID] = p
	fmt.Println("Producto agregado correctamente")
}

func (i *Inventario) EliminarProducto(id int) {
	if _, existe := i.Productos[id]; existe {
		delete(i.Productos, id)
		fmt.Println("Producto eliminado correctamente")
	} else {
		fmt.Println("Producto no encontrado")
	}
}

func (i *Inventario) BuscarProducto(id int) (Producto, bool) {
	producto, existe := i.Productos[id]
	return producto, existe
}

func (i *Inventario) ListarProductos() {
	if len(i.Productos) == 0 {
		fmt.Println("El inventario esta vacio")
		return
	}

	fmt.Println("\n=== LISTA DE PRODUCTOS ===")
	for _, producto := range i.Productos {
		producto.MostrarInfo()
	}
}

func main() {
	inventario := NuevoInventario()
	var opcion int

	for {
		fmt.Println("\n===== MENU DE INVENTARIO =====")
		fmt.Println("1. Agregar producto")
		fmt.Println("2. Eliminar producto")
		fmt.Println("3. Buscar producto")
		fmt.Println("4. Actualizar precio")
		fmt.Println("5. Actualizar cantidad")
		fmt.Println("6. Listar productos")
		fmt.Println("7. Salir")
		fmt.Print("Seleccione una opcion: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			var id, cantidad int
			var nombre string
			var precio float64

			fmt.Print("Ingrese ID: ")
			fmt.Scan(&id)

			fmt.Print("Ingrese nombre: ")
			fmt.Scan(&nombre)

			fmt.Print("Ingrese precio: ")
			fmt.Scan(&precio)

			fmt.Print("Ingrese cantidad: ")
			fmt.Scan(&cantidad)

			producto := NuevoProducto(id, nombre, precio, cantidad)
			inventario.AgregarProducto(producto)

		case 2:
			var id int

			fmt.Print("Ingrese ID del producto a eliminar: ")
			fmt.Scan(&id)

			inventario.EliminarProducto(id)

		case 3:
			var id int

			fmt.Print("Ingrese ID del producto a buscar: ")
			fmt.Scan(&id)

			producto, existe := inventario.BuscarProducto(id)

			if existe {
				producto.MostrarInfo()
			} else {
				fmt.Println("Producto no encontrado")
			}

		case 4:
			var id int
			var nuevoPrecio float64

			fmt.Print("Ingrese ID del producto: ")
			fmt.Scan(&id)

			producto, existe := inventario.BuscarProducto(id)

			if existe {
				fmt.Print("Ingrese nuevo precio: ")
				fmt.Scan(&nuevoPrecio)

				producto.ActualizarPrecio(nuevoPrecio)
				inventario.Productos[id] = producto

				fmt.Println("Precio actualizado correctamente")
			} else {
				fmt.Println("Producto no encontrado")
			}

		case 5:
			var id, nuevaCantidad int

			fmt.Print("Ingrese ID del producto: ")
			fmt.Scan(&id)

			producto, existe := inventario.BuscarProducto(id)

			if existe {
				fmt.Print("Ingrese nueva cantidad: ")
				fmt.Scan(&nuevaCantidad)

				producto.ActualizarCantidad(nuevaCantidad)
				inventario.Productos[id] = producto

				fmt.Println("Cantidad actualizada correctamente")
			} else {
				fmt.Println("Producto no encontrado")
			}

		case 6:
			inventario.ListarProductos()

		case 7:
			fmt.Println("Saliendo del sistema...")
			return

		default:
			fmt.Println("Opcion no valida")
		}
	}
}

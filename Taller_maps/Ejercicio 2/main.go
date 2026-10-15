package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func calcularTotal(p Producto) float64 {
	return p.Precio * float64(p.Cantidad)
}

func tieneStock(p Producto) bool {
	return p.Cantidad > 0
}

func mostrarProducto(p Producto) {
	fmt.Println("-------------------------")
	fmt.Println("Nombre:", p.Nombre)
	fmt.Printf("Precio: $%.2f\n", p.Precio)
	fmt.Println("Cantidad:", p.Cantidad)
	fmt.Printf("Valor total: $%.2f\n", calcularTotal(p))

	if tieneStock(p) {
		fmt.Println("Estado: Disponible")
	} else {
		fmt.Println("Estado: Sin stock")
	}
}

func main() {
	var productos [3]Producto

	fmt.Println("=== REGISTRO DE PRODUCTOS ===")

	for i := 0; i < 3; i++ {
		fmt.Printf("\nProducto %d\n", i+1)

		fmt.Print("Ingrese el nombre: ")
		fmt.Scan(&productos[i].Nombre)

		fmt.Print("Ingrese el precio: ")
		fmt.Scan(&productos[i].Precio)

		fmt.Print("Ingrese la cantidad: ")
		fmt.Scan(&productos[i].Cantidad)
	}

	fmt.Println("\n=== LISTADO DE PRODUCTOS ===")

	for i := 0; i < 3; i++ {
		mostrarProducto(productos[i])
	}
}

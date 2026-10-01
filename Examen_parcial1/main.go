package main

import "fmt"

var productosVendidos []string
var subtotales []float64

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)

	productosVendidos = append(productosVendidos, nombre)
	subtotales = append(subtotales, subtotal)

	fmt.Printf("Venta registrada: %s - Cantidad: %d - Subtotal: $%.2f\n",
		nombre, cantidad, subtotal)
}

func MostrarEstadisticas() {
	if len(subtotales) == 0 {
		fmt.Println("No existen ventas registradas.")
		return
	}

	total := 0.0

	for _, subtotal := range subtotales {
		total += subtotal
	}

	fmt.Println("\n--- ESTADÍSTICAS ---")
	fmt.Println("Total recaudado: $", total)
	fmt.Println("Número de ventas registradas:", len(subtotales))
}

func main() {

	nombres := []string{"Arroz", "Leche", "Pan"}
	precios := []float64{1.25, 0.95, 0.50}

	var opcion int

	for {
		fmt.Println("\n===== MENÚ =====")
		fmt.Println("1. Registrar una nueva venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		switch opcion {

		case 1:
			var productoSeleccionado int
			var cantidad int

			fmt.Println("\n--- PRODUCTOS DISPONIBLES ---")
			fmt.Println("1. Arroz - $1.25")
			fmt.Println("2. Leche - $0.95")
			fmt.Println("3. Pan   - $0.50")

			fmt.Print("Seleccione el número del producto: ")
			fmt.Scan(&productoSeleccionado)

			if productoSeleccionado < 1 || productoSeleccionado > 3 {
				fmt.Println("Producto no válido.")
				continue
			}

			fmt.Print("Ingrese la cantidad vendida: ")
			fmt.Scan(&cantidad)

			if cantidad <= 0 {
				fmt.Println("La cantidad debe ser mayor que 0.")
				continue
			}

			indice := productoSeleccionado - 1

			RegistrarVenta(
				nombres[indice],
				precios[indice],
				cantidad,
			)

		case 2:
			MostrarEstadisticas()

		case 3:
			fmt.Println("Programa finalizado.")
			return

		default:
			fmt.Println("Opción no válida.")
		}
	}
}

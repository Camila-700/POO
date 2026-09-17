package main

import "fmt"

func saludar() {
	fmt.Println("Hola esta es mi primera funcion")
}

// Funcion con un parametro
func nombre(nombre string) {
	fmt.Println("Bienvenida/a", nombre)
}
func main() {
	var usr string
	fmt.Println("Ingresa tu nombre:")
	fmt.Scan(&usr)
	fmt.Println("Hola bienvenido/a:", usr)
}

// solicite dos numeros y resulva su suma
func suma(a int, b int) int {
	return a + b
}

func main() {
	var usr string
	fmt.Println("Ingresa tu nombre:")
	fmt.Scan(&usr)
	saludar()
	bienvenida(usr)

	fmt.Println("El resutado de la suma es : ", suma(4, 5))
}

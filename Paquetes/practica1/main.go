package main

import (
	"fmt"
	"practica/operaciones"
	"practica/saludo"
)

func main() {
	fmt.Println("Bienvenidos a las clase de Paquetes😁")
	mensaje := saludo.Saludar("Camila")
	fmt.Println(mensaje)

	fmt.Println(operaciones.Suma(5, 6))
}

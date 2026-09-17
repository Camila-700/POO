package main

import "fmt"

/*
func <nombre>(param1, param2, ....param n)<valores de retorno>{
---------------------------
-------------------------------
---------------------------------

		//return en el caso de que nuestra funcion retorne valores
	}
*/
func saludar() {
	fmt.Println("Hola esta es mi primera función")

}
func Bienvenida(nombre string) {
	fmt.Println("Bienvenid@", nombre)

}
func suma(a, b int) int {
	return a + b
}
func sumaresta(c, d int) (int, int) {
	return c + d, c - d
}
func main() {
	var usr string

	fmt.Println("Ingresa tu nombre")
	fmt.Scan(&usr)
	saludar()
	Bienvenida(usr)

	var a, b int
	fmt.Println("Ingresa dos valores:  ")
	fmt.Scan(&a, &b)
	fmt.Println("El resultado de la suma es: ", suma(a, b))

	var c, d int
	fmt.Println("Ingrese dos valores, para relaizar una suma y resta: ")
	fmt.Scan(&c, &d)

	w, z := sumaresta(c, d)
	fmt.Println("El resultado de la suma es: ", w, " y el resultado de la resta es:  ", z)

}

package main

import (
	"fmt"
	contador_vocales "taller/Contador_vocales"
	conversor_monedas "taller/Conversor_monedas"
)

func main() {

	var opcion int

	fmt.Println("1. Conversor de monedas")
	fmt.Println("2. Contador de vocales")
	fmt.Print("Ingrese una opcion: ")
	fmt.Scanln(&opcion)

	if opcion == 1 {

		var dolares float64
		var moneda string

		fmt.Print("Ingrese el valor en dolares: ")
		fmt.Scanln(&dolares)

		fmt.Println("Monedas permitidas:")
		fmt.Println("Euros")
		fmt.Println("LB")
		fmt.Println("Won")
		fmt.Println("BTC")

		fmt.Print("Ingrese la moneda: ")
		fmt.Scanln(&moneda)

		resultado := conversor_monedas.Convertir(dolares, moneda)

		fmt.Println("El resultado es:", resultado)

	} else if opcion == 2 {

		var frase string

		fmt.Print("Ingrese una frase: ")
		fmt.Scanln(&frase)

		a, e, i, o, u := contador_vocales.ContarVocales(frase)

		fmt.Println("Cantidad de vocales:")
		fmt.Println("a =", a)
		fmt.Println("e =", e)
		fmt.Println("i =", i)
		fmt.Println("o =", o)
		fmt.Println("u =", u)

	} else {

		fmt.Println("Opcion incorrecta")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== CONTEO DE PALABRAS ===")
	fmt.Print("Ingrese una frase: ")
	texto, _ := reader.ReadString('\n')

	texto = strings.ToLower(texto)
	palabras := strings.Fields(texto)

	conteo := make(map[string]int)

	for _, palabra := range palabras {
		conteo[palabra]++
	}

	fmt.Println("\nRepeticiones de palabras:")

	palabraMayor := ""
	mayor := 0

	for palabra, cantidad := range conteo {
		fmt.Printf("%s: %d\n", palabra, cantidad)

		if cantidad > mayor {
			mayor = cantidad
			palabraMayor = palabra
		}
	}

	if mayor > 0 {
		fmt.Println("\nPalabra mas repetida:", palabraMayor)
		fmt.Println("Repeticiones:", mayor)
	} else {
		fmt.Println("No se ingresaron palabras.")
	}
}

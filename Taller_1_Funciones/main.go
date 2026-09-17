package main

import "fmt"

func averageGrade(suma float64, cantidad int) float64 {
	return suma / float64(cantidad)
}

func estudiantes() {
	var cantidad, nota int
	var total float64

	fmt.Println("Ingrese la cantidad de estudiantes:")
	fmt.Scan(&cantidad)

	for i := 1; i <= cantidad; i++ {
		fmt.Println("Ingrese la nota del estudiante", i)
		fmt.Scan(&nota)

		for nota < 0 || nota > 100 {
			fmt.Println("Error: la nota debe estar entre 0 y 100")
			fmt.Println("Ingrese nuevamente la nota:")
			fmt.Scan(&nota)
		}

		total += float64(nota)
	}
	promedio := averageGrade(total, cantidad)

	fmt.Println("El promedio del curso es:", promedio)

	if promedio >= 70 {
		fmt.Println("Curso aprobado")
	} else {
		fmt.Println("Curso reprobado")
	}

	switch {
	case promedio >= 90:
		fmt.Println("Excellent performance")
	case promedio >= 80:
		fmt.Println("Good performance")
	case promedio >= 70:
		fmt.Println("Satisfactory performance")
	default:
		fmt.Println("Needs improvement")
	}
}

func sumatoria() {
	var n, total int

	fmt.Println("Ingrese un numero:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		total += i
	}

	fmt.Println("La suma de los numeros del 1 al", n, "es:", total)
}
func celsiustoFahrenheit() {
	var celsius, fahrenheit float64

	fmt.Println("Ingrese la temperatura en Celsius:")
	fmt.Scan(&celsius)

	fahrenheit = (celsius * 9 / 5) + 32

	fmt.Println("La temperatura en Fahrenheit es:", fahrenheit)
}

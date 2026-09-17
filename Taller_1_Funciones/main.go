package main

import "fmt"

func averageGrade(suma float64, cantidad int) float64 {
	return suma / float64(cantidad)
}

func estudiantes() {
	var cantidad, nota, total int

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

		total += nota
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

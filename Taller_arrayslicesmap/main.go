package main

import "fmt"

// EJERCICIO 1: ANÁLISIS DE NOTAS DE ESTUDIANTES

func ejercicio1() {
	fmt.Println("====================================")
	fmt.Println("EJERCICIO 1: NOTAS DE ESTUDIANTES")
	fmt.Println("====================================")

	notas := [6][4]float64{
		{8.5, 9.0, 7.5, 10.0},
		{7.0, 8.0, 9.0, 8.5},
		{9.5, 9.0, 10.0, 9.0},
		{6.5, 7.0, 8.0, 7.5},
		{8.0, 8.5, 9.0, 9.5},
		{7.5, 8.0, 7.0, 8.5},
	}

	promedios := []float64{}
	notasAltas := []float64{}
	notasBajas := []float64{}

	sumaGeneral := 0.0

	for _, estudiante := range notas {
		suma := 0.0
		mayor := estudiante[0]
		menor := estudiante[0]

		for _, nota := range estudiante {
			suma += nota

			if nota > mayor {
				mayor = nota
			}

			if nota < menor {
				menor = nota
			}
		}

		promedio := suma / float64(len(estudiante))

		promedios = append(promedios, promedio)
		notasAltas = append(notasAltas, mayor)
		notasBajas = append(notasBajas, menor)

		sumaGeneral += suma
	}

	for i := 0; i < len(promedios); i++ {
		fmt.Printf("\nEstudiante %d\n", i+1)
		fmt.Printf("Promedio: %.2f\n", promedios[i])
		fmt.Printf("Nota más alta: %.2f\n", notasAltas[i])
		fmt.Printf("Nota más baja: %.2f\n", notasBajas[i])
	}

	promedioGeneral := sumaGeneral / float64(6*4)

	fmt.Println("\n------------------------------------")
	fmt.Printf("Promedio general de la clase: %.2f\n", promedioGeneral)
	fmt.Println("------------------------------------")
}

// EJERCICIO 2: VOTACIÓN DE ACTIVIDADES CON MAPS

func actividadGanadora(votos map[string]int) string {
	ganadora := ""
	mayor := -1

	for actividad, cantidad := range votos {
		if cantidad > mayor {
			mayor = cantidad
			ganadora = actividad
		}
	}

	return ganadora
}

func ejercicio2() {
	fmt.Println("\n====================================")
	fmt.Println("EJERCICIO 2: VOTACIÓN DE ACTIVIDADES")
	fmt.Println("====================================")

	votos := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	var actividad string

	fmt.Println("\nActividades disponibles:")
	fmt.Println("1. deportes")
	fmt.Println("2. videojuegos")
	fmt.Println("3. cine")
	fmt.Println("4. musica")

	for i := 1; i <= 5; {
		fmt.Printf("\nIngrese el voto %d: ", i)
		fmt.Scan(&actividad)

		if _, existe := votos[actividad]; existe {
			votos[actividad]++
			fmt.Println("Voto registrado correctamente.")
			i++
		} else {
			fmt.Println("Actividad no válida. Intente nuevamente.")
		}
	}

	fmt.Println("\n------------------------------------")
	fmt.Println("RESULTADOS DE LA VOTACIÓN")
	fmt.Println("------------------------------------")

	for actividad, cantidad := range votos {
		fmt.Printf("%s: %d votos\n", actividad, cantidad)
	}

	ganadora := actividadGanadora(votos)

	fmt.Println("\n------------------------------------")
	fmt.Printf("Actividad ganadora: %s\n", ganadora)
	fmt.Printf("Cantidad de votos: %d\n", votos[ganadora])
	fmt.Println("------------------------------------")
}

// FUNCIÓN PRINCIPAL

func main() {
	// Ejecutar el ejercicio 1
	ejercicio1()

	// Ejecutar el ejercicio 2
	ejercicio2()
}

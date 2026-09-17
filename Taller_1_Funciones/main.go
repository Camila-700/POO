package main

func averageGrade(suma float64, cantidad int) float64 {
	return suma / float64(cantidad)
}
func estudiantes() {
	var cantidad int
	var nota, total float64

	fmt.Println("Ingrese la cantidad de estudiantes:")
	fmt.Scan(&cantidad)


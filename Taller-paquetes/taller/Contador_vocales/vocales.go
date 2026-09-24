package contador_vocales

func ContarVocales(frase string) (int, int, int, int, int) {

	var a, e, i, o, u int

	for _, letra := range frase {

		switch letra {

		case 'a', 'A':
			a++

		case 'e', 'E':
			e++

		case 'i', 'I':
			i++

		case 'o', 'O':
			o++

		case 'u', 'U':
			u++
		}
	}

	return a, e, i, o, u
}

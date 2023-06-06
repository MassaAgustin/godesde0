package ejercicios

import "strconv"

func ConversionNumerica(valor string) (int, string) {
	conversion, err := strconv.Atoi(valor)

	if err != nil {
		return 0, "Hubo un error en la conversión" + err.Error()
	}

	if conversion > 100 {
		return conversion, "El valor es mayor a 100"
	} else {
		return conversion, "El valor es menor a 100"
	}
}

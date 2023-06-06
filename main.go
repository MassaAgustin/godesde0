package main

import (
	"fmt"
	"godesde0/ejercicios"
)

func main() {
	entero, cadena := ejercicios.ConversionNumerica("fff")

	fmt.Println("El valor es: ", entero)
	fmt.Println("El texto es: ", cadena)
}

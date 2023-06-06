package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var err error
var numero1 int

func TablaDeMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un numero: ")

		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(numero1, " * ", i, " = ", numero1*i)
	}
}

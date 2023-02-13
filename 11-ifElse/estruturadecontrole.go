package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controles")

	numero := 10

	if numero > 5 {
		fmt.Println("numero é maior que 5")
	} else {
		fmt.Println("numero é menor que 5")
	}

	teste := false

	if VouF := teste; VouF != true { // if init - fica limitada a esse escopo, se tentar chamar ela fora do escopo ela não funciona
		fmt.Println("O valor é:", teste)
	} else { // else if pode ser usado também
		fmt.Println("O valor é:", teste)
	}

}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i:", i)
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println("Resultado: ", i)

	// for j := 0; j <= 10; j++ {
	// 	fmt.Println("Incrementando j: ", j)
	// 	time.Sleep(time.Second)
	// 	if j == 10 {
	// 		fmt.Println("Resultado: ", j)
	// 	}
	// }

	// nomes := [3]string{"João", "Davi", "Lucas"}

	// for _, nome := range nomes {
	// 	fmt.Println(nome)
	// }

	// for indice, letra := range "PALAVRA" {
	// 	fmt.Println(indice, string(letra))
	// }

	usuario := map[string]string{
		"nome":      "José",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}

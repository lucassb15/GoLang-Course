package main

import "fmt"

func main() {

	retorno := func(texto string) string { //func anonima
		return fmt.Sprintf("Recebendo -> %s %d", texto, 1) // Consegue passar uma string e concatenar com tipos de dados
	}("Passando o valor do Parâmetro")
	fmt.Println(retorno)

}

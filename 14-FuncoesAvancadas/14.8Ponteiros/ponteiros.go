package main

import "fmt"

func inverterSinal(numero int) int { // parâmetro por cópia
	return numero * -1
}

// quando quer alterar o valor da variavel utiliza-se o ponteiro
func inverterSinalPonteiro(numero *int) { // passando uma referência
	*numero = *numero * -1
	// desreferenciação *numero > ele vai buscar o valor no endereço de memória
	// referênciando *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero) // não altera o valor da variavel !

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalPonteiro(&novoNumero) // & para jogar o endereço de memoria na  funcão
	fmt.Println(novoNumero)            // agora o valor da variavel é alterado !

}

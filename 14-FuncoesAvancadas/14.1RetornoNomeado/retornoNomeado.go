package main

import "fmt"

//func com retorno nomeado
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
	// return soma, subtracao  X
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma)
	fmt.Println(subtracao)
}

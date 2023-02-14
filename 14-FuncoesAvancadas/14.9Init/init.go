package main

import "fmt"

var n int

func init() { // criar setups
	fmt.Println("Executando a função init")
	n = 10
}

func main() { // pode ter uma função por arquivo, mas não pode utilizar mais que uma por pacote
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}

package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Println("Métodos")
	usuario1 := usuario{"Lucas", 23}
	fmt.Println(usuario1)
	// usuario1.salvarDados() precisa de um método
}

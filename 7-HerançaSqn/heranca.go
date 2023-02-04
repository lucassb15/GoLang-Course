package main

import "fmt"

type pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Altura    uint8
}

type estudante struct {
	pessoa    // herança = só utilizar o nome do outro struct
	Curso     string
	Faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{
		"Lucas",
		"Barbosa",
		22,
		179,
	}
	fmt.Println(p1)

	e1 := estudante{
		p1,
		"x",
		"y",
	}
	fmt.Println(e1)
}

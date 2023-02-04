package main

import "fmt"

type usuario struct {
	Nome     string
	Idade    uint8
	endereco endereco
}

type endereco struct {
	Logradouro string
	Numero     uint8
}

func main() {
	fmt.Println("Arquvio structs")

	var u usuario
	u.Nome = "Lucas"
	u.Idade = 22
	fmt.Println(u)

	enderecoExemplo := endereco{
		Logradouro: "Rua",
		Numero:     5,
	}

	usuario2 := usuario{
		"Lucas",
		21,
		enderecoExemplo,
	}
	fmt.Println(usuario2)

	usuario3 := usuario{
		Nome:  "Lucas",
		Idade: 0,
	}
	fmt.Println(usuario3)
}

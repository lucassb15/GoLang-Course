package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvado os dados do Usuário %s no banco de dados \n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")
	usuario1 := usuario{"Lucas", 17}
	fmt.Println(usuario1)
	// usuario1.salvarDados() precisa de um método
	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade())
	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()
	fmt.Println(usuario2.maiorDeIdade())

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}

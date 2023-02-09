package main

import "fmt"

func main() {
	fmt.Println("Maps")
	// [string] = tipo das chaves  // string = tipos de dados, eles são fixos, todos são do tipo string ou todos são do tipo int, não pode mesclar tipos de dados
	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Barbosa",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"]) // acessar um campo especifico

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Lucas",
			"ultimo":   "Barbosa",
		},
		"curso": {
			"nome":   "ADS",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
}

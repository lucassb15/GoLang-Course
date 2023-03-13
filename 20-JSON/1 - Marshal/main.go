package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

// conversão de struct para JSON
func main() {
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Toddy",
		"raca": "Poodle",
	}
	cachorro2EmJSON, erro := json.Marshal(c2)
	fmt.Println()
	fmt.Println()
	fmt.Println(c2)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))

}

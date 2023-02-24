package main

import (
	"fmt"
	"log"
	"modulos/app"
	"os"
)

func main() {
	fmt.Println("oi")
	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}

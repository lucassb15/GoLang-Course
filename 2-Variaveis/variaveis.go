package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	varivel2 := "Variável 2"
	fmt.Println(variavel1, "  ", varivel2)

	var (
		variavel3 string = "123"
		variavel4 string = "1234"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Var5", "Var6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}

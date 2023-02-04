package main

import "fmt"

func main() {
	//Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	somas := numero1 + numero2 // usar mesmos tipos  int16,16 ok, int16, 32 fail
	fmt.Println(somas)
	// Fim dos aritméticos

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// Fim dos operadores de atribuição

	// Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// Fim dos relacionais

	// Operadores Lógicos
	verdadeiro, falso := true, false
	fmt.Println("--------------")
	fmt.Println(verdadeiro && falso) // verifica se todas são verdadeira
	fmt.Println(verdadeiro || falso) // verifica as condições se uma for verdadeira retorna true
	fmt.Println(!verdadeiro)         // inverter o valor
	fmt.Println(!falso)              // inverter o valor
	// fim operadores lógicos

	// Operadores Unários
	numero := 10
	// numero = numero + 1
	numero++     // um em um
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20
	fmt.Println(numero)

	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)
	// Fim dos operadores unários

	// texto := numero > 5 ? "Maior que 5" : "Menor que 5" não funciona
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}

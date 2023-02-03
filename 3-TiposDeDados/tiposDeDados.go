package main

import "fmt"

func main() {
	var numero int16 = 10000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE / ascii
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)
	// Fim Números Inteiros

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1234569123.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)
	// Fim Números Reias

	var str string = "abcdefgh"
	fmt.Println(str)

	str2 := "abcdefghijk"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)
	// Fim Strings

	var texto string
	fmt.Println(texto)
}

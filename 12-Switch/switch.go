package main

import "fmt"

func diaDaSemana(numero int) string {
	// Utilizando if else
	// if numero == 1 {
	// 	return "Domingo"
	// } else if numero == 2 {
	// 	return "Segunda feira"
	// }

	// com Switch
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número invalido"
	}

} // Aponta um erro falando que precisa de um return no final da função

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda feira"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número invalido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(2)
	fmt.Println(dia)
	dia = diaDaSemana(3)
	fmt.Println(dia)
	dia = diaDaSemana(4)
	fmt.Println(dia)
	dia = diaDaSemana(5)
	fmt.Println(dia)
	fmt.Println("---")
	fmt.Println("Var")
	dia2 := diaDaSemana(6)
	fmt.Println(dia2)
	dia2 = diaDaSemana(7)
	fmt.Println(dia2)
	dia2 = diaDaSemana(8)
	fmt.Println(dia2)
	dia2 = diaDaSemana(9)
	fmt.Println(dia2)
}

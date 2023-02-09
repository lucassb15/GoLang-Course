package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Pos 1" //popular
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"} // usando inferência de tipos
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	// Fim arrays

	// Slice
	slice := []int{11, 15, 16, 10, 5, 6, 7, 8, 9} // ele não é um array, mas parece com array
	fmt.Println(slice)

	slice = append(slice, 35)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2) //parece um ponteiro pegando no array2

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	// Arrays internos
	slice3 := make([]float32, 10, 11) // tamanho : 10 , capacidade: 15, se estourar a capacidade o Go dobra a capacidade
	fmt.Println(slice3)
	fmt.Println("tamanho: ", len(slice3))    // length
	fmt.Println("capacidade: ", cap(slice3)) // capacidade
	fmt.Println("---------")
	fmt.Println("---------")
	// estourar a capacidade
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println("tamanho: ", len(slice3))    // length
	fmt.Println("capacidade: ", cap(slice3)) // capacidade
	fmt.Println("---------")
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}

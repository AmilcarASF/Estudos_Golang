package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"Teste1", "Teste2", "Teste3", "Teste4", "Test5"}
	array2[1] = "Posição 2"
	fmt.Println(array2)

	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array2))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 20)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "alterado"
	fmt.Println(array2)
	fmt.Println(slice2)

	// arrays internos
	fmt.Println("-------")

	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 50)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}

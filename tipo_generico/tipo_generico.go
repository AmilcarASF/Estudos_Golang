package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Tipo Genérico")

	generica("string")
	generica(100)
	generica(true)
}

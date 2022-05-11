package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
}

func funcao2() {
	fmt.Println("Função 2")
}

func main() {
	fmt.Println("Defer")

	defer funcao1()
	funcao2()

}

package main

import "fmt"

func somar(num1, num2 int8) int8 {
	return num1 + num2
}

func soma_sub(num1, num2 int8) (int8, int8) {
	soma := num1 + num2
	sub := num1 - num2
	return soma, sub
}

func main() {
	fmt.Println(somar(10, 20))
	soma, sub := soma_sub(50, 5)
	fmt.Println(soma, sub)

	soma2, _ := soma_sub(100, 5)
	fmt.Println(soma2)
}

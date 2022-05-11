package main

import "fmt"

func diaDaSemana(numero int) string {
	var dia string

	switch numero {
	case 1:
		dia = "Domingo"
		fallthrough
	case 2:
		dia = "Segunda-Feira"
	case 3:
		dia = "Terça-Feira"
	case 4:
		dia = "Quarta-Feira"
	case 5:
		dia = "Quinta-Feira"
	case 6:
		dia = "Sexta-Feira"
	case 7:
		dia = "Sábado"
	default:
		dia = "Erro"
	}

	return dia
}

func main() {

	if num1 := 10; num1 > 0 {
		fmt.Println("Maior que zero")
		fmt.Println(num1)
	}

	fmt.Println(diaDaSemana(1))

}

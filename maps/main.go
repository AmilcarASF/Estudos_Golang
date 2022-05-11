package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	usuario["descr"] = "Fisica"

	fmt.Println(usuario)

	delete(usuario, "sobrenome")

	fmt.Println(usuario)

}

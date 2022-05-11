package main

import (
	"errors"
	"fmt"
)

func main() {

	var erro error = errors.New("not implemented")
	fmt.Println(erro)

}

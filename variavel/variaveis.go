package main

import "fmt"

func main() {
	var var1 string = "var 1"
	var2 := "var 2"
	fmt.Println(var1, var2)

	var (
		var3 string = "teste"
		var4 string = "test"
	)

	fmt.Println(var3, var4)

	var5, var6 := "var 5", "var 6"

	fmt.Println(var5, var6)

	var5, var6 = var6, var5
	fmt.Println(var5, var6)

}

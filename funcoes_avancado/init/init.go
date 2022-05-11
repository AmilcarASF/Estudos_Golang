package main

import "fmt"

var n int

func init() {
	fmt.Println("exec func init")
	n = 10
}

func main() {
	fmt.Println("init")
	fmt.Println(n)

}

package main

import "fmt"

func main() {
	fmt.Println(Add(2, 3))
	fmt.Println(Sub(3, 2))
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

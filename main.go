package main

import "fmt"

func main() {
	fmt.Println(Add(2, 3))
	fmt.Println(Sub(3, 2))
	fmt.Println(Mul(3, 2))
	fmt.Println(Div(3, 2))
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

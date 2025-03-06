package main

import "fmt"

func add(num1 int, num2 int) int{
	sum := num1 + num2

	return sum
}

func main(){

	a := 10
	b := 20

	sum := a + b

	fmt.Println("Sum: ", sum)
	summation := add(a, b)

	fmt.Println("Sum with function: ", summation)
}
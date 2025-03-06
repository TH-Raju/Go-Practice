package main

import "fmt"

func add(num1 int, num2 int) int{
	sum := num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2

	return sum, mul

}

func main(){

	a := 10
	b := 20

	sum := a + b

	fmt.Println("Sum: ", sum)
	summation := add(a, b)

	fmt.Println("Sum with function: ", summation)

	p, q := getNumbers(a, b)
	fmt.Println("p: ", p, " q: ", q)
}
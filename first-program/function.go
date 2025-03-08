package main

import "fmt"

func add(num1 int, num2 int) int{
	sum := num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int, int) {
	sum := num1 + num2
	mul := num1 * num2
	div := num1 / num2

	return sum, mul, div
}

func printSomewhere(){
	fmt.Println("Education must be free")
}

func sayHello(name string) {
	fmt.Println("Welcome to the golang ", name)
}

func main(){

	a := 10
	b := 20

	sum := a + b

	fmt.Println("Sum: ", sum)
	summation := add(a, b)

	fmt.Println("Sum with function: ", summation)
	printSomewhere()

	sayHello("Raju")

	p, q, r := getNumbers(a, b)
	fmt.Println("p: ", p, " q: ", q, "r: ", r)
}
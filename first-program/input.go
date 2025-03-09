package main

import "fmt"


func main(){
	fmt.Println("Starting")

	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Printf("Hello, %s", name)

	var num1 int
	var num2 int
	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)

	sum := num1 + num2

	// Display results

	fmt.Println("Hello, ", name)
	fmt.Println("Summation: ", sum)

	// Goodbye message

	fmt.Println("Thank your for your Information")
	fmt.Println("Goodbye, ", name)
}
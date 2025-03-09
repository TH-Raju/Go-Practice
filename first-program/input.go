package main

import "fmt"


func printWelcomeMessage(){
	fmt.Println("Welcome to the Application")
}

func getUserName() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func getTwoNumber() (int, int){
	fmt.Println("Enter first number: ")
    var num1 int
    fmt.Scanln(&num1)

    fmt.Println("Enter second number: ")
    var num2 int
    fmt.Scanln(&num2)

    return num1, num2  
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func displayResults(name string, sum int){
	fmt.Println("Hello, ", name)
	fmt.Println("Summation: ", sum)
}

func getGoodbyeMessage(name string) {
    fmt.Println("Thank your for your Information")
	fmt.Println("Goodbye, ", name)
}

func main(){
	printWelcomeMessage()
	name := getUserName()

	num1, num2 := getTwoNumber()
	sum := add(num1, num2)
	
	// Display results
	displayResults(name, sum)
	
	// Goodbye message
	getGoodbyeMessage(name)
	
}
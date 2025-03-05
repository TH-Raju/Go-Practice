package main

import "fmt"

func add(num1 int, num2 int){
	fmt.Println("Sum of ", num1 , "+",num2, "=" , num1 + num2)
}

func main(){

	a := 10
	b := 20

	sum := a + b

	fmt.Println("Sum: ", sum)
	add(a, b)
}
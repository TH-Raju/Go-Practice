package main

import "fmt"


func main(){
	fmt.Println("Starting")

	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Printf("Hello, %s", name)
}
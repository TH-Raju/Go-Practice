package main

import "fmt" // fmt = format

func main(){
	a := 1 // variable
	/*
	int
	float32
	bool
	string
	*/
	fmt.Println(a)

	var b = 2.5


	fmt.Println("Hello on Go in", b,"st Ramadan!");


	age := 21
	if age >= 21 {
		fmt.Println("You are eligible to drive!")
	}else if age >= 18 {
		fmt.Println("You are not eligible to drive!")
	}
}



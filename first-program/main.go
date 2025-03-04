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


	age := 2
	if age >= 21 {
		fmt.Println("You are eligible to drive!")
	}else if age >= 18 {
		fmt.Println("You are not eligible to drive!")
	}else {
		fmt.Println("You are not old enough to drive!")
	}


	x := 3

	switch x {
		case 1:
            fmt.Println("One")
        case 2, 3:
            fmt.Println("Two or Three")
        default:
            fmt.Println("Unknown")
	}
}



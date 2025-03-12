package main
import "fmt"
var (
	a = 20
	b = 30
)

/*
	There are 3 types of Scope
	1. Global Scope
	2. Local Scope
	3. Package Scope
*/

func add(x int, y int){
	z := x + y
	fmt.Println(z)
}

func main(){
	fmt.Println("Hello, Scope")
	p := 30
	q := 40

	add(p, q)

	add(a, b)

	add(a, p)

	if p == 30{
		fmt.Println("p is ", a)
	}

	switch
}
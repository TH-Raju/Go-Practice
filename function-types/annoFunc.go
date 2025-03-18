package main
import "fmt"


var name = "Raju"

// Standard or named Function
func add(a,b int){
	fmt.Println(a+b)
}


func main() {
	// we invoke the add function here
	add(2,4)


	// anonymous function 
	// Immediately Invoke Function Expression
	// IIFE
	func (a int, b int){
		c := a + b
		fmt.Println(c)
	}(5,10) // (5,10) this is for call to function

}

func init(){
	fmt.Println("Hey",name ," This is Init function and it's executed before main function")
}
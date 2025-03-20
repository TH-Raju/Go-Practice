package main
import "fmt"


var a = 10
func main(){
	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scan(&name)

	// fmt.Println("Hello, ", name)

	fmt.Println("Hello at first start Init Function then main function called...")
	fmt.Println(a)

}

func init(){
	fmt.Println("I am the first function that execute first...")
	fmt.Println("From init",a)
	a = 20
}
package main
import "fmt"


func main(){
	add := func(a int, b int){
		c := a+b
		fmt.Println(c)
	}

	add(2,3)
}

func init(){
	fmt.Println("I'll be call first")
}
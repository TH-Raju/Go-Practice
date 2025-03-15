package main
import ("fmt" 
"example.com/mathlib"
"example.com/test")

func main(){
	fmt.Println("Learn Package Scope!!")
	mathlib.Add(4,6)
	test.Hello()
}

// go mod init example.com  



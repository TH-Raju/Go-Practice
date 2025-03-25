package main
import "fmt"


func main(){
	var name string
	var ans string
	fmt.Println("Hello!!")
	fmt.Print("May I know your name? : ")
	fmt.Scan(&name)
	fmt.Println("Hello!", name, "Are you ready to start the Quiz......!!!")
	fmt.Print("Enter (Y) for yes and (N) for no: ")
	fmt.Scan(&ans)

	if(ans == "Y" || ans == "y"){
		fmt.Println("Let's begin the quiz!")
	}else {
		fmt.Println("Okay, Quiz is over!")
        return
	}



}
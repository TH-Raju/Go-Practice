package main
import "fmt"

var score int

func question( ques string, correctAns int) {
	var userAns int
	fmt.Print("Write Answer:")
	fmt.Scan(&userAns)

	if(userAns == correctAns){
		score++
		fmt.Println("Correct Answer!")
	}else {
		fmt.Println("Wrong Answer!")
	}

}

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

		question("1 + 1 = ? ", 2)
		question("4 * 4 = ? ", 16)
		question("In binary 1 + 1 = ? ", 10)
		

		fmt.Println("Your Score : ", score)
	}else {
		fmt.Println("Okay, Quiz is over!")
        return
	}



}
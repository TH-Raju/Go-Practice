package main
import "fmt"

var score int

func intQuestion( ques string, correctAns int) {
	var userAns int
	fmt.Println(ques)
	fmt.Print("Write Answer:")
	fmt.Scan(&userAns)

	if(userAns == correctAns){
		score++
		fmt.Println("Correct Answer!")
	}else {
		fmt.Println("Wrong Answer!")
	}

}

func intro(){
	var name string
	fmt.Println("Hello!!")
	fmt.Print("May I know your name? : ")
	fmt.Scan(&name)
	fmt.Println("Hello!", name, "Are you ready to start the Quiz......!!!")
}

func main(){
	var ans string
	
	intro()
	fmt.Print("Enter (Y) for yes and (N) for no: ")
	fmt.Scan(&ans)

	if(ans == "Y" || ans == "y"){
		fmt.Println("Let's begin the quiz!")

		intQuestion("1 + 1 = ? ", 2)
		intQuestion("4 * 4 = ? ", 16)
		intQuestion("In binary 1 + 1 = ? ", 10)
		
	}
	fmt.Println("Okay, Quiz is over!")
	fmt.Println("Your Score : ", score ,"\n\n")


	fmt.Println("Thanks for your time")
      

}
package main
import ("fmt"
"strings")

var score int
var quesNo int

func intQuestion( ques string, correctAns int) {
	quesNo++
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

func textQuestion(ques string, correctAns string){
	quesNo++
	var userAns string
	fmt.Println(ques)
	fmt.Print("Write Answer: ")
	fmt.Scan(&userAns)

	if strings.ToLower(userAns) == strings.ToLower(correctAns) {
		score++
		fmt.Println("Correct Answer!")
	}else{
		fmt.Println("Wrong Answer!!")
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

	if(strings.ToLower(ans) == "y"){
		fmt.Println("Let's begin the quiz!")

		// Integer based questions
		intQuestion("1 + 1 = ? ", 2)
		intQuestion("4 * 4 = ? ", 16)
		intQuestion("In binary 1 + 1 = ? ", 10)
		
		// Text based questions
		textQuestion("What is the capital of Bangladesh? ", "dhaka")
		textQuestion("Who is creator ", "Allah")
	}
	fmt.Println("Okay, Quiz is over!")
	fmt.Println("Your Score : ", score, "/", quesNo ,"\n\n")


	fmt.Println("Thanks for your time")
      

}